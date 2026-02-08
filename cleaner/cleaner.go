package cleaner

import (
	"os"
	"path/filepath"
	"sync"
	"wincu/utils"
	"wincu/worker"
)

// Cleaner handles the cleaning process
type Cleaner struct {
	Targets []Target
	DryRun  bool
	Force   bool
	Pool    *worker.Pool
}

// NewCleaner creates a new Cleaner instance
func NewCleaner(targets []Target, dryRun bool, force bool, workerCount int) *Cleaner {
	return &Cleaner{
		Targets: targets,
		DryRun:  dryRun,
		Force:   force,
		Pool:    worker.NewPool(workerCount),
	}
}

// Run executes the cleaning process
func (c *Cleaner) Run() {
	c.Pool.Start()
	defer c.Pool.Stop()

	var wg sync.WaitGroup

	for _, target := range c.Targets {
		if !target.Enabled {
			continue
		}

		if target.RequireAdmin && !utils.IsAdmin() {
			utils.Warn("Skipping " + target.Name + " (Requires Admin)")
			continue
		}

		utils.Info("Scanning " + target.Name + " (" + target.Path + ")...")

		// Process each target
		wg.Add(1)
		// We process targets sequentially to walk usage, but files concurrently?
		// Actually, walking can be slow. We can walk in parallel?
		// For now, let's walk in main goroutine (or separate goroutine per target) and submit file deletion tasks.
		// Doing it per target in goroutine is better.

		go func(t Target) {
			defer wg.Done()
			c.processTarget(t)
		}(target)
	}

	wg.Wait()
}

func (c *Cleaner) processTarget(target Target) {
	if target.Name == "Recycle Bin" {
		// Submit as a single task
		c.Pool.Submit(func() {
			EmptyRecycleBin(c.DryRun)
		})
		return
	}

	// If path doesn't exist, skip
	if _, err := os.Stat(target.Path); os.IsNotExist(err) {
		utils.Warn("Target path not found: " + target.Path)
		return
	}

	err := filepath.Walk(target.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Permission denied or other error
			// utils.Debug("Error accessing path: " + path + " - " + err.Error())
			return nil // Continue walking
		}

		if path == target.Path {
			return nil // Don't delete root
		}

		if info.IsDir() {
			// We might want to delete empty directories later, but requirement says "Delete files only" in Safety Rules?
			// "228: ลบเฉพาะ file เท่านั้น" -> "Delete files only"
			// But usually temp folders accumulate empty dirs too.
			// Let's stick to files for now as per "Safety Rules".
			return nil
		}

		// Submit file deletion task
		c.Pool.Submit(func() {
			DeleteFile(path, c.DryRun, c.Force)
		})

		return nil
	})

	if err != nil {
		utils.Error("Error walking target " + target.Name + ": " + err.Error())
	}
}
