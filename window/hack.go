package window

// #include <SFML/Graphics.h>
import "C"

import (
	"unsafe"

	"github.com/mewspring/sfml/font"
	"github.com/mewspring/sfml/texture"
)

// drawableHack is a copy of texture.Drawable without modifications. Through the
// use of unsafe and with knowledge of its memory layout we are able to access
// unexported members. This hack allows us to cross package barriers while
// keeping the exported API clean.
type drawableHack struct {
	// A renderable GPU texture.
	tex *C.sfRenderTexture
	// A sprite representation of the GPU texture.
	sprite *C.sfSprite
}

// drawableSprite returns the sprite of the provided texture.Drawable.
func drawableSprite(tex texture.Drawable) *C.sfSprite {
	return (*drawableHack)(unsafe.Pointer(&tex)).sprite
}

// imageHack is a copy of texture.Image without modifications. Through the use
// of unsafe and with knowledge of its memory layout we are able to access
// unexported members. This hack allows us to cross package barriers while
// keeping the exported API clean.
type imageHack struct {
	// A read-only GPU texture.
	tex *C.sfTexture
	// A sprite representation of the GPU texture.
	sprite *C.sfSprite
}

// imageSprite returns the sprite of the provided texture.Image.
func imageSprite(tex texture.Image) *C.sfSprite {
	return (*imageHack)(unsafe.Pointer(&tex)).sprite
}

// textHack is a copy of font.Text without modifications. Through the use of
// unsafe and with knowledge of its memory layout we are able to access
// unexported members. This hack allows us to cross package barriers while
// keeping the exported API clean.
type textHack struct {
	// A graphical text entry.
	text *C.sfText
}

// textText returns the text of the provided font.Text.
func textText(text font.Text) *C.sfText {
	return (*textHack)(unsafe.Pointer(&text)).text
}
