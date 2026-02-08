package cleaner

import (
	"os"
	"path/filepath"
	"wincu/utils"
)

// ScanResult holds the result of a scan
type ScanResult struct {
	Target string
	Size   int64
	Count  int
}

// ScanTargets scans the given targets and returns the total size and file count
func ScanTargets(targets []Target) []ScanResult {
	var results []ScanResult

	for _, target := range targets {
		if !target.Enabled {
			continue
		}

		// If Recycle Bin, uses a different generic way or just skip size calculation for now?
		// Recycle Bin size calculation requires SHQueryRecycleBin.
		// For now, we skip size for Recycle Bin or mock it.
		if target.Name == "Recycle Bin" {
			// TODO: Implement Recycle Bin size check
			results = append(results, ScanResult{Target: target.Name, Size: 0, Count: 0})
			continue
		}

		if _, err := os.Stat(target.Path); os.IsNotExist(err) {
			continue
		}

		if target.RequireAdmin && !utils.IsAdmin() {
			utils.Warn("Skipping scan of " + target.Name + " (Requires Admin)")
			continue
		}

		var size int64
		var count int

		err := filepath.Walk(target.Path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if !info.IsDir() {
				size += info.Size()
				count++
			}
			return nil
		})

		if err != nil {
			utils.Error("Error scanning " + target.Name + ": " + err.Error())
		}

		results = append(results, ScanResult{Target: target.Name, Size: size, Count: count})
	}
	return results
}
