// Package window handles window creation, drawing and events.
//
// This package uses a small subset of the features provided by the SFML library
// version 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package window

// #include <SFML/Graphics/RenderWindow.h>
import "C"

import (
	"image"
	"image/color"

	"github.com/mewmew/wandi"
)

// A sfmlWindow represents a graphical window capable of handling draw
// operations and window events. It implements the wandi.Window interface.
type sfmlWindow struct {
	w *C.sfRenderWindow
}

// Open opens a window with the specified dimensions.
//
// Note: The Close method of the window must be called when finished using it.
func Open(width, height int) (win wandi.Window, err error) {
	sfmlWin := new(sfmlWindow)
	mode := C.sfVideoMode{
		width:        C.uint(width),
		height:       C.uint(height),
		bitsPerPixel: 32,
	}
	title := C.CString("untitled")
	style := C.sfUint32(0)
	sfmlWin.w = C.sfRenderWindow_create(mode, title, style, nil)
	return sfmlWin, nil
}

// Close closes the window.
func (sfmlWin *sfmlWindow) Close() {
	C.sfRenderWindow_close(sfmlWin.w)
	C.sfRenderWindow_destroy(sfmlWin.w)
}

// SetTitle sets the title of the window.
//
// Note: The title will be updated on the next call to PollEvent.
func (sfmlWin *sfmlWindow) SetTitle(title string) {
	C.sfRenderWindow_setUnicodeTitle(sfmlWin.w, utf32(title))
}

// Update displays window updates on the screen.
func (sfmlWin *sfmlWindow) Update() (err error) {
	C.sfRenderWindow_display(sfmlWin.w)
	return nil
}

// Clear clears the screen and fills it with the provided color.
func (sfmlWin *sfmlWindow) Clear(c color.Color) (err error) {
	C.sfRenderWindow_clear(sfmlWin.w, sfmlColor(c))
	return nil
}

// Draw draws the entire src image onto the window starting at the destination
// point dp.
func (sfmlWin *sfmlWindow) Draw(dp image.Point, src wandi.Image) (err error) {
	panic("sfmlWindow.Draw: not yet implemented.")
}

// DrawRect fills the destination rectangle dr of the window with corresponding
// pixels from the src image starting at the source point sp.
func (sfmlWin *sfmlWindow) DrawRect(dr image.Rectangle, src wandi.Image, sp image.Point) (err error) {
	panic("sfmlWindow.DrawRect: not yet implemented.")
}
