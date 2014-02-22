// Package window handles window creation, drawing and events.
//
// This package uses a small subset of the features provided by the SFML library
// version 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package window

// #include <SFML/Graphics.h>
import "C"

import (
	"fmt"
	"image"
	"image/color"

	"github.com/mewmew/sfml/font"
	"github.com/mewmew/sfml/texture"
	"github.com/mewmew/wandi"
)

// Style specifies the style and behavior of windows.
type Style int

// Window styles.
const (
	// Fixes states that the window cannot be resized.
	Fixed Style = iota
	// FullScreen states that the window is in full screen mode.
	FullScreen
)

// A Window represents a graphical window capable of handling draw operations
// and window events. It implements the wandi.Window interface.
type Window struct {
	// A renderable window.
	Win *C.sfRenderWindow
}

// Open opens a window with the specified dimensions and an optional window
// style. By default the window is resizeable and not in full screen mode.
//
// Note: The main thread must be used for both window creation and event
// handling. It is perfectly fine to use separate threads for rendering and
// event handling, as long as all event handling takes place in the main thread.
//
// Note: The Close method of the window must be called when finished using it.
func Open(width, height int, style ...Style) (win wandi.Window, err error) {
	sfWin := new(Window)
	mode := C.sfVideoMode{
		width:        C.uint(width),
		height:       C.uint(height),
		bitsPerPixel: 32,
	}
	title := C.CString("untitled")
	sfStyle := C.sfUint32(C.sfDefaultStyle)
	if len(style) > 0 {
		switch style[0] {
		case Fixed:
			sfStyle &^= C.sfResize
		case FullScreen:
			sfStyle = C.sfFullscreen
		}
	}
	sfWin.Win = C.sfRenderWindow_create(mode, title, sfStyle, nil)

	// TODO(u): Enable vsync.
	//C.sfRenderWindow_setVerticalSyncEnabled(sfWin.Win, C.sfTrue)

	// TODO(u): Deactivate OpenGL context for the current thread. This way
	// whichever thread that ends up being the rendering thread will activate the
	// OpenGL context implicitly through a call to window.display().
	//C.sfRenderWindow_setActive(C.sfFalse)

	return sfWin, nil
}

// Close closes the window.
func (win *Window) Close() {
	C.sfRenderWindow_close(win.Win)
	C.sfRenderWindow_destroy(win.Win)
}

// SetTitle sets the title of the window.
//
// Note: The title will be updated on the next call to PollEvent.
func (win *Window) SetTitle(title string) {
	C.sfRenderWindow_setUnicodeTitle(win.Win, utf32(title))
}

// Width returns the width of the window.
func (win *Window) Width() int {
	size := C.sfWindow_getSize(win.Win)
	return int(size.x)
}

// Height returns the height of the window.
func (win *Window) Height() int {
	size := C.sfWindow_getSize(win.Win)
	return int(size.y)
}

// Clear clears the window and fills it with the provided color.
func (win *Window) Clear(c color.Color) {
	C.sfRenderWindow_clear(win.Win, sfmlColor(c))
}

// Draw draws the entire src image onto the window starting at the destination
// point dp.
func (win *Window) Draw(dp image.Point, src wandi.Image) (err error) {
	sr := image.Rect(0, 0, src.Width(), src.Height())
	return win.DrawRect(dp, src, sr)
}

// DrawRect draws a subset of the src image, as defined by the source rectangle
// sr, onto the window starting at the destination point dp.
func (win *Window) DrawRect(dp image.Point, src wandi.Image, sr image.Rectangle) (err error) {
	switch srcImg := src.(type) {
	case *font.Text:
		// TODO(u): Handle sr?
		C.sfText_setPosition(srcImg.Text, sfmlFloatPt(dp))
		C.sfRenderWindow_drawText(win.Win, srcImg.Text, nil)
	case *texture.Texture:
		C.sfSprite_setTextureRect(srcImg.Sprite, sfmlIntRect(sr))
		C.sfSprite_setPosition(srcImg.Sprite, sfmlFloatPt(dp))
		C.sfRenderWindow_drawSprite(win.Win, srcImg.Sprite, nil)
	default:
		return fmt.Errorf("Window.DrawRect: support for image format %T not yet implemented", src)
	}
	return nil
}

// Update displays window rendering updates on the screen.
func (win *Window) Update() {
	C.sfRenderWindow_display(win.Win)
}
