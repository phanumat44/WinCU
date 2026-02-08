package cleaner

import (
	"os"
	"wincu/utils"
)

// DeleteFile attempts to delete a file and logs the result
func DeleteFile(path string, dryRun bool, force bool) error {
	if dryRun {
		utils.Info("Would delete: " + path)
		return nil // Simulate success
	}

	err := os.Remove(path)
	if err != nil {
		// If force is enabled and it's a permission error, try to remove read-only attribute
		if force && os.IsPermission(err) {
			// Try to change mode to 0666 (rw-rw-rw-)
			chmodErr := os.Chmod(path, 0666)
			if chmodErr == nil {
				// Retry deletion
				err = os.Remove(path)
			}
		}

		if err != nil {
			if os.IsPermission(err) {
				utils.Warn("Permission denied: " + path)
			} else {
				utils.Warn("Failed to delete: " + path + " - " + err.Error())
			}
			return err
		}
	}

	utils.Info("Deleted: " + path)
	return nil
}
