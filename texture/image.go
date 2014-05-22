package texture

// #cgo LDFLAGS: -lcsfml-graphics
// #include <SFML/Graphics.h>
import "C"

import (
	"errors"
	"fmt"
	"image"
)

// Image represent a read-only texture. It implements the wandi.Image interface.
type Image struct {
	// A read-only GPU texture.
	tex *C.sfTexture
	// A sprite representation of the GPU texture.
	sprite *C.sfSprite
}

// Load loads the provided file and converts it into a read-only texture.
//
// Note: The Free method of the texture must be called when finished using it.
func Load(path string) (tex Image, err error) {
	// Load texture from file.
	tex.tex = C.sfTexture_createFromFile(C.CString(path), nil)
	if tex.tex == nil {
		return Image{}, fmt.Errorf("texture.Load: unable to load %q", path)
	}

	// Create a sprite for the texture.
	tex.sprite = C.sfSprite_create()
	if tex.sprite == nil {
		return Image{}, errors.New("texture.Load: unable to create sprite")
	}
	C.sfSprite_setTexture(tex.sprite, tex.tex, C.sfTrue)

	return tex, nil
}

// Read reads the provided image and converts it into a read-only texture.
func Read(img image.Image) (tex Image, err error) {
	panic("not yet implemented")
}

// Free frees the texture.
func (tex Image) Free() {
	// TODO(u): Free sprite as well?
	C.sfTexture_destroy(tex.tex)
}

// Width returns the width of the texture.
func (tex Image) Width() int {
	size := C.sfTexture_getSize(tex.tex)
	return int(size.x)
}

// Height returns the height of the texture.
func (tex Image) Height() int {
	size := C.sfTexture_getSize(tex.tex)
	return int(size.y)
}
