// +build linux

package netreuse

import (
	"syscall"
)

var soReuseAddr = syscall.SO_REUSEADDR
