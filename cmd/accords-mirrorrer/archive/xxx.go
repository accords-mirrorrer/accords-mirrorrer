//go:build !windows

package archive

import (
	"syscall"
)

var (
	SIGUSR1 = syscall.SIGUSR1
)