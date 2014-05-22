// Package window handles window creation, drawing and events.
//
// This package uses a small subset of the features provided by the SFML library
// version 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package window

// #cgo LDFLAGS: -lcsfml-graphics
// #include <SFML/Graphics.h>
import "C"

import (
	"image"
	"image/color"

	"github.com/mewmew/wandi"
)

// Style specifies the style and behavior of windows.
type Style int

// Window styles.
const (
	// Fixed states that the window cannot be resized.
	Fixed Style = iota
	// FullScreen states that the window is in full screen mode.
	FullScreen
)

// A Window represents a graphical window capable of handling draw operations
// and window events. It implements the wandi.Window interface.
type Window struct {
	// A renderable window.
	win *C.sfRenderWindow
}

// Open opens a window of the specified dimensions and with an optional window
// style. By default the window is resizeable and not in full screen mode.
func Open(width, height int, style ...Style) {
	panic("not yet implemented")
}

// Close closes the window.
func (win Window) Close() {
	panic("not yet implemented")
}

// SetTitle sets the title of the window.
func (win Window) SetTitle(title string) {
	panic("not yet implemented")
}

// Width returns the width of the window.
func (win Window) Width() int {
	panic("not yet implemented")
}

// Height returns the height of the window.
func (win Window) Height() int {
	panic("not yet implemented")
}

// Draw draws the entire src image onto the window starting at the destination
// point dp.
func (win Window) Draw(dp image.Point, src wandi.Image) (err error) {
	panic("not yet implemented")
}

// DrawRect draws a subset of the src image, as defined by the source rectangle
// sr, onto the window starting at the destination point dp.
func (win Window) DrawRect(dp image.Point, src wandi.Image, sr image.Rectangle) (err error) {
	panic("not yet implemented")
}

// Fill fills the entire window with the provided color.
func (win Window) Fill(c color.Color) {
	panic("not yet implemented")
}

// Display displays what has been rendered so far to the window.
func (win Window) Display() {
	panic("not yet implemented")
}
