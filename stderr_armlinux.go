//go:build (linux && arm) || (linux && arm64)

package utils

import (
	"os"
	"syscall"
)

func RedirectStdErrToFile(f *os.File) error {
	return syscall.Dup3(int(f.Fd()), int(os.Stderr.Fd()), 0) // 将stderr重定向到文件，比如捕获panic
}
