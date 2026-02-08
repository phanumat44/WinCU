package utils

import (
	"syscall"
)

// IsAdmin checks if the current process utilizes administrator privileges.
func IsAdmin() bool {
	shell32 := syscall.NewLazyDLL("shell32.dll")
	isUserAnAdmin := shell32.NewProc("IsUserAnAdmin")

	ret, _, _ := isUserAnAdmin.Call()
	return ret != 0
}
