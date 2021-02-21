// +build darwin freebsd dragonfly netbsd openbsd linux

package netreuse

import (
	"sync"
	"sync/atomic"
)

// checker is a struct to gather the availability check fields + funcs.
// we use atomic ints because this is potentially a really hot function call.
type checkerT struct {
	avail int32      // atomic int managed by set/isAvailable()
	check int32      // atomic int managed by has/checked()
	mu    sync.Mutex // synchonizes the actual check
}

// the static location of the vars.
var checker checkerT

func (c *checkerT) isAvailable() bool {
	return atomic.LoadInt32(&c.avail) != 0
}

func (c *checkerT) setIsAvailable(b bool) {
	if b {
		atomic.StoreInt32(&c.avail, 1)
	} else {
		atomic.StoreInt32(&c.avail, 0)
	}
}

func (c *checkerT) hasChecked() bool {
	return atomic.LoadInt32(&c.check) != 0
}

func (c *checkerT) setHasChecked(b bool) {
	if b {
		atomic.StoreInt32(&c.check, 1)
	} else {
		atomic.StoreInt32(&c.check, 0)
	}
}

// Available returns whether or not SO_REUSEPORT is available in the OS.
// It does so by attepting to open a tcp listener, setting the option, and
// checking ENOPROTOOPT on error. After checking, the decision is cached
// for the rest of the process run.
func available() bool {
	checker.setIsAvailable(true)
	checker.setHasChecked(true)
	return checker.isAvailable()
}
