// +build linux

package main

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"

func init() {
	// ref: http://en.sfml-dev.org/forums/index.php?topic=10228.msg100775
	C.XInitThreads()
}
