// Package window handles window creation, drawing and events. It uses a small
// subset of the features provided by the SFML library version 2.5 [1].
//
// [1]: http://www.sfml-dev.org/
package window

// #include <SFML/Graphics.h>
//
// #cgo LDFLAGS: -lcsfml-graphics -lcsfml-window
import "C"

import (
	"fmt"
	"image"
	"image/color"

	"github.com/mewspring/sfml/font"
	"github.com/mewspring/sfml/texture"
	"github.com/mewspring/wandi"
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

// Open opens a new window of the specified dimensions. An optional window style
// may be provided. By default the window is resizeable and not in full screen
// mode.
//
// Note: The main thread must be used for both window creation and event
// handling. It is perfectly fine to use separate threads for rendering and
// event handling, as long as all event handling takes place in the main thread.
//
// Note: The Close method of the window must be called when finished using it.
func Open(width, height int, style ...Style) (*Window, error) {
	if len(style) > 1 {
		return nil, fmt.Errorf("window.Open: invalid number of optional window styles; expected zero or one, got %d", len(style))
	}

	// Open a new window of the specified dimensions.
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
	w := C.sfRenderWindow_create(mode, title, sfStyle, nil)
	win := &Window{
		win: w,
	}

	// TODO(u): Decide if vsync should be enabled by default.

	// Enable vsync.
	//C.sfRenderWindow_setVerticalSyncEnabled(sfWin.Win, C.sfTrue)

	// TODO(u): Evaluate the effect of deactivating the OpenGL context of the
	// current thread.

	// Deactivate the OpenGL context of the current thread. Whichever thread that
	// ends up being the rendering thread will activate the OpenGL context
	// implicitly through a call to window.display().
	//C.sfRenderWindow_setActive(C.sfFalse)

	return win, nil
}

// Close closes the window.
func (win *Window) Close() {
	C.sfRenderWindow_close(win.win)
	C.sfRenderWindow_destroy(win.win)
}

// SetTitle sets the title of the window.
//
// Note: The title will be updated on the next call to PollEvent.
func (win *Window) SetTitle(title string) {
	C.sfRenderWindow_setUnicodeTitle(win.win, utf32(title))
}

// ShowCursor displays or hides the mouse cursor depending on the value of
// visible. It is visible by default.
func (win *Window) ShowCursor(visible bool) {
	C.sfRenderWindow_setMouseCursorVisible(win.win, sfmlBool(visible))
}

// Width returns the width of the window.
func (win *Window) Width() int {
	size := C.sfRenderWindow_getSize(win.win)
	return int(size.x)
}

// Height returns the height of the window.
func (win *Window) Height() int {
	size := C.sfRenderWindow_getSize(win.win)
	return int(size.y)
}

// Draw draws the entire src image onto the window starting at the destination
// point dp.
func (win *Window) Draw(dp image.Point, src wandi.Image) error {
	sr := image.Rect(0, 0, src.Width(), src.Height())
	return win.DrawRect(dp, src, sr)
}

// DrawRect draws a subset of the src image, as defined by the source rectangle
// sr, onto the window starting at the destination point dp.
func (win *Window) DrawRect(dp image.Point, src wandi.Image, sr image.Rectangle) error {
	switch srcImg := src.(type) {
	case *texture.Drawable:
		sprite := drawableSprite(srcImg)
		C.sfSprite_setTextureRect(sprite, sfmlIntRect(sr))
		C.sfSprite_setPosition(sprite, sfmlFloatPt(dp))
		C.sfRenderWindow_drawSprite(win.win, sprite, nil)
	case *texture.Image:
		sprite := imageSprite(srcImg)
		C.sfSprite_setTextureRect(sprite, sfmlIntRect(sr))
		C.sfSprite_setPosition(sprite, sfmlFloatPt(dp))
		C.sfRenderWindow_drawSprite(win.win, sprite, nil)
	case *font.Text:
		// TODO(u): Handle sr?
		text := textText(srcImg)
		C.sfText_setPosition(text, sfmlFloatPt(dp))
		C.sfRenderWindow_drawText(win.win, text, nil)
	default:
		return fmt.Errorf("Window.DrawRect: support for image format %T not yet implemented", src)
	}

	return nil
}

// Fill fills the entire window with the provided color.
func (win *Window) Fill(c color.Color) {
	C.sfRenderWindow_clear(win.win, sfmlColor(c))
}

// SetActive activates the CPU context of the window.
func (win *Window) SetActive() {
	C.sfRenderWindow_setActive(win.win, C.sfTrue)
}

// Display displays what has been rendered so far to the window.
func (win *Window) Display() {
	C.sfRenderWindow_display(win.win)
}
