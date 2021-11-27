package texture

// #include <SFML/Graphics.h>
//
// #cgo LDFLAGS: -lcsfml-graphics
import "C"

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"log"
	"time"
	"unsafe"
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
func Load(path string) (*Image, error) {
	// Load the texture from file.
	t := C.sfTexture_createFromFile(C.CString(path), nil)
	if t == nil {
		return nil, fmt.Errorf("texture.Load: unable to load %q", path)
	}
	tex := &Image{
		tex: t,
	}
	// Create a sprite for the texture.
	sprite := C.sfSprite_create()
	if sprite == nil {
		return nil, errors.New("texture.Load: unable to create sprite")
	}
	tex.sprite = sprite
	C.sfSprite_setTexture(tex.sprite, tex.tex, C.sfTrue)
	return tex, nil
}

// Read reads the provided image and converts it into a read-only texture.
//
// Note: The Free method of the texture must be called when finished using it.
func Read(src image.Image) (*Image, error) {
	// Use fallback conversion for unknown image formats.
	rgba, ok := src.(*image.RGBA)
	if !ok {
		return Read(fallback(src))
	}
	// Use fallback conversion for subimages.
	width, height := rgba.Rect.Dx(), rgba.Rect.Dy()
	if rgba.Stride != 4*width {
		return Read(fallback(src))
	}
	// Create a read-only texture based on the pixels of the src image.
	t := C.sfTexture_create(C.uint(width), C.uint(height))
	if t == nil {
		return nil, fmt.Errorf("texture.Read: unable to create %dx%d texture", width, height)
	}
	tex := &Image{
		tex: t,
	}
	pix := (*C.sfUint8)(unsafe.Pointer(&rgba.Pix[0]))
	C.sfTexture_updateFromPixels(tex.tex, pix, C.uint(width), C.uint(height), 0, 0)
	// Create a sprite for the texture.
	sprite := C.sfSprite_create()
	if sprite == nil {
		return nil, errors.New("texture.Read: unable to create sprite")
	}
	tex.sprite = sprite
	C.sfSprite_setTexture(tex.sprite, tex.tex, C.sfTrue)
	return tex, nil
}

// fallback converts the provided image or subimage into a RGBA image.
func fallback(src image.Image) *image.RGBA {
	start := time.Now()

	// Create a new RGBA image and draw the src image onto it.
	bounds := src.Bounds()
	dr := image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	dst := image.NewRGBA(dr)
	draw.Draw(dst, dr, src, bounds.Min, draw.Src)

	log.Printf("texture.fallback: fallback conversion for non-RGBA image (%T) finished in: %v", src, time.Since(start))

	return dst
}

// Free frees the texture.
func (tex *Image) Free() {
	C.sfSprite_destroy(tex.sprite)
	C.sfTexture_destroy(tex.tex)
}

// Width returns the width of the texture.
func (tex *Image) Width() int {
	size := C.sfTexture_getSize(tex.tex)
	return int(size.x)
}

// Height returns the height of the texture.
func (tex *Image) Height() int {
	size := C.sfTexture_getSize(tex.tex)
	return int(size.y)
}
