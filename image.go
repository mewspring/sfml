// Package sfml handles image drawing operations.
//
// This package uses a small subset of the features provided by the SFML library
// version 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package sfml

import (
	"image"

	"github.com/mewmew/wandi"
)

// A sfmlImage is a mutable collection of pixels, which is stored in CPU memory.
// It implements the wandi.Image interface.
type sfmlImage struct {
}

// TODO(u): Add note about free to NewImage, LoadImage and ReadImage?
//
// Note: The Free method of the image should be called when finished using it.

// NewImage returns a new image of the specified dimensions. The image is stored
// in CPU memory.
func NewImage(width, height int) (img wandi.Image, err error) {
	panic("sfml.NewImage: not yet implemented.")
}

// LoadImage loads the provided image file and returns it as an image. The image
// is stored in CPU memory.
func LoadImage(imgPath string) (img wandi.Image, err error) {
	panic("sfml.LoadImage: not yet implemented.")
}

// ReadImage reads the provided image, converts it to the standard image format
// of this library and returns it. The image is stored in CPU memory.
func ReadImage(src image.Image) (img wandi.Image, err error) {
	panic("sfml.ReadImage: not yet implemented.")
}

// Free frees the image.
func (img *sfmlImage) Free() {
	panic("sfmlImage.Free: not yet implemented.")
}

// Draw draws the entire src image onto the dst image starting at the
// destination point dp.
func (dst *sfmlImage) Draw(dp image.Point, src wandi.Image) (err error) {
	panic("sfmlImage.Draw: not yet implemented.")
}

// DrawRect fills the destination rectangle dr of the dst image with
// corresponding pixels from the src image starting at the source point sp.
func (dst *sfmlImage) DrawRect(dr image.Rectangle, src wandi.Image, sp image.Point) (err error) {
	panic("sfmlImage.DrawRect: not yet implemented.")
}
