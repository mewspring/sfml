package texture

// #cgo LDFLAGS: -lcsfml-graphics
// #include <SFML/Graphics.h>
import "C"

import (
	"image"
)

// Image represent a read-only texture. It implements the wandi.Image interface.
type Image struct {
	// A read-only GPU texture.
	tex *C.sfTexture
}

// Load loads the provided file and converts it into a read-only texture.
func Load(path string) (tex Image, err error) {
	panic("not yet implemented")
}

// Read reads the provided image and converts it into a read-only texture.
func Read(img image.Image) (tex Image, err error) {
	panic("not yet implemented")
}

// Free frees the texture.
func (tex Image) Free() {
	panic("not yet implemented")
}

// Width returns the width of the texture.
func (tex Image) Width() int {
	panic("not yet implemented")
}

// Height returns the height of the texture.
func (tex Image) Height() int {
	panic("not yet implemented")
}
