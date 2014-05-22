package window

// #include <SFML/Graphics.h>
import "C"

import (
	"unsafe"

	"github.com/mewmew/sfml/texture"
)

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

func imageSprite(tex texture.Image) *C.sfSprite {
	return (*imageHack)(unsafe.Pointer(&tex)).sprite
}
