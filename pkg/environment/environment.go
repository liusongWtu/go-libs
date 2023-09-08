package environment

import (
	"os"
	"runtime"
)

const (
	SysWindows = "windows"
	SysMac     = "darwin"
)

func IsWindows() bool {
	return runtime.GOOS == SysWindows
}

func IsMac() bool {
	return runtime.GOOS == SysMac
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
