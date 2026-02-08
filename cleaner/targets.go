package cleaner

import (
	"os"
	"path/filepath"
)

// Target represents a cleanup target
type Target struct {
	Name         string
	Path         string
	RequireAdmin bool
	Enabled      bool
}

// GetTargets returns the list of cleanup targets
func GetTargets() []Target {
	userProfile := os.Getenv("USERPROFILE")
	if userProfile == "" {
		// Fallback or handle error
		userProfile = "C:\\Users\\Default"
	}

	return []Target{
		{
			Name:         "User Temp",
			Path:         filepath.Join(userProfile, "AppData", "Local", "Temp"),
			RequireAdmin: false,
			Enabled:      true,
		},
		{
			Name:         "Windows Temp",
			Path:         "C:\\Windows\\Temp",
			RequireAdmin: true,
			Enabled:      true,
		},
		{
			Name:         "Prefetch",
			Path:         "C:\\Windows\\Prefetch",
			RequireAdmin: true,
			Enabled:      true,
		},
		{
			Name:         "Windows Update Cache",
			Path:         "C:\\Windows\\SoftwareDistribution\\Download",
			RequireAdmin: true,
			Enabled:      true,
		},
		{
			Name:         "Recycle Bin",
			Path:         "Windows API",
			RequireAdmin: false, // Standard user can empty own recycle bin usually, but let's check.
			Enabled:      true,
		},
	}
}
