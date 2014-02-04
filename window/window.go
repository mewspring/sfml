// Package win handles window creation, drawing and events.
//
// This package uses a small subset of the features provided by the SFML library
// version 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package win

import (
	"image"
	"image/color"

	"github.com/mewmew/wandi"
)

// A sfmlWindow represents a graphical window capable of handling draw
// operations and window events. It implements the wandi.Window interface.
type sfmlWindow struct {
}

// Open opens a window with the specified dimensions.
//
// Note: The Close method of the window must be called when finished using it.
func Open(width, height int) (win wandi.Window, err error) {
	panic("win.Open: not yet implemented")

	sfmlWin := new(sfmlWindow)
	return sfmlWin, nil
}

// Close closes the window.
func (win *sfmlWindow) Close() {
	panic("sfmlWindow.Close: not yet implemented.")
}

// SetTitle sets the title of the window.
func (win *sfmlWindow) SetTitle(title string) {
	panic("sfmlWindow.SetTitle: not yet implemented.")
}

// Update displays window updates on the screen.
func (win *sfmlWindow) Update() (err error) {
	panic("sfmlWindow.Update: not yet implemented.")
}

// Clear clears the screen and fills it with the provided color.
func (win *sfmlWindow) Clear(c color.Color) (err error) {
	panic("sfmlWindow.Clear: not yet implemented.")
}

// Draw draws the entire src image onto the window starting at the destination
// point dp.
func (win *sfmlWindow) Draw(dp image.Point, src wandi.Image) (err error) {
	panic("sfmlWindow.Draw: not yet implemented.")
}

// DrawRect fills the destination rectangle dr of the window with corresponding
// pixels from the src image starting at the source point sp.
func (win *sfmlWindow) DrawRect(dr image.Rectangle, src wandi.Image, sp image.Point) (err error) {
	panic("sfmlWindow.DrawRect: not yet implemented.")
}
