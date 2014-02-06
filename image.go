// Package sfml handles image drawing operations.
//
// This package uses a small subset of the features provided by the SFML library
// version 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package sfml

// #include <SFML/Graphics.h>
import "C"

import (
	"fmt"
	"image"

	"github.com/mewmew/wandi"
)

// An Image is a mutable collection of pixels, which is stored in CPU memory.
// It implements the wandi.Image interface.
type Image struct {
	// Img is a SFML image stored in CPU memory.
	Img *C.sfImage
}

// NewImage returns a new image of the specified dimensions. The image is stored
// in CPU memory.
//
// Note: The Free method of the image should be called when finished using it.
func NewImage(width, height int) (img wandi.Image, err error) {
	sfmlImg := new(Image)
	sfmlImg.Img = C.sfImage_create(C.uint(width), C.uint(height))
	if sfmlImg.Img == nil {
		return nil, fmt.Errorf("sfml.NewImage: unable to create %dx%d image", width, height)
	}
	return sfmlImg, nil
}

// LoadImage loads the provided image file and returns it as an image. The image
// is stored in CPU memory.
//
// Note: The Free method of the image should be called when finished using it.
func LoadImage(filePath string) (img wandi.Image, err error) {
	sfmlImg := new(Image)
	sfmlImg.Img = C.sfImage_createFromFile(C.CString(filePath))
	if sfmlImg.Img == nil {
		return nil, fmt.Errorf("sfml.LoadImage: unable to load image %q", filePath)
	}
	return sfmlImg, nil
}

// ReadImage reads the provided image, converts it to the standard image format
// of this library and returns it. The image is stored in CPU memory.
//
// Note: The Free method of the image should be called when finished using it.
func ReadImage(src image.Image) (img wandi.Image, err error) {
	panic("sfml.ReadImage: not yet implemented.")
}

// Free frees the image.
func (img *Image) Free() {
	C.sfImage_destroy(img.Img)
}

// Width returns the width of the image.
func (img *Image) Width() int {
	size := C.sfImage_getSize(img.Img)
	return int(size.x)
}

// Height returns the height of the image.
func (img *Image) Height() int {
	size := C.sfImage_getSize(img.Img)
	return int(size.y)
}

// Draw draws the entire src image onto the dst image starting at the
// destination point dp.
func (dst *Image) Draw(dp image.Point, src wandi.Image) (err error) {
	sr := image.Rect(0, 0, src.Width(), src.Height())
	return dst.DrawRect(dp, src, sr)
}

// DrawRect fills the destination rectangle dr of the dst image with
// corresponding pixels from the src image starting at the source point sp.
func (dst *Image) DrawRect(dp image.Point, src wandi.Image, sr image.Rectangle) (err error) {
	switch srcImg := src.(type) {
	case *Image:
		srcRect := sfmlIntRect(sr)
		C.sfImage_copyImage(dst.Img, srcImg.Img, C.uint(dp.X), C.uint(dp.Y), srcRect, C.sfTrue)
	default:
		return fmt.Errorf("Image.DrawRect: support for image format %T not yet implemented", src)
	}
	return nil
}
