package cleaner

import (
	"syscall"
	"wincu/utils"
)

var (
	modShell32             = syscall.NewLazyDLL("shell32.dll")
	procSHEmptyRecycleBinW = modShell32.NewProc("SHEmptyRecycleBinW")
)

const (
	SHERB_NOCONFIRMATION = 0x00000001
	SHERB_NOPROGRESSUI   = 0x00000002
	SHERB_NOSOUND        = 0x00000004
)

// EmptyRecycleBin empties the Recycle Bin
func EmptyRecycleBin(dryRun bool) error {
	if dryRun {
		utils.Info("Would empty Recycle Bin")
		return nil
	}

	utils.Info("Emptying Recycle Bin...")

	// SHEmptyRecycleBinW(HWND hwnd, LPCTSTR pszRootPath, DWORD dwFlags)
	// hwnd = 0 (no window)
	// pszRootPath = 0 (all drives)
	// dwFlags = NOCONFIRMATION | NOPROGRESSUI | NOSOUND

	ret, _, _ := procSHEmptyRecycleBinW.Call(
		0,
		0,
		uintptr(SHERB_NOCONFIRMATION|SHERB_NOPROGRESSUI|SHERB_NOSOUND),
	)

	if ret != 0 {
		utils.Error("Failed to empty Recycle Bin. Error code: %d", ret)
		return syscall.Errno(ret)
	}

	utils.Info("Recycle Bin emptied successfully")
	return nil
}
