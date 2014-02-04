// Package sfml handles image drawing operations using the SFML library version
// 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package sfml

import (
	"image"

	"github.com/mewmew/wandi"
)

// A sfmlImg is a mutable collection of pixels. It implements the wandi.Image
// interface.
type sfmlImg struct {
}

// TODO(u): Add note about free to NewImage, LoadImage and ReadImage?
//
// Note: The Free method of the image should be called when finished using it.

// NewImage returns a new image of the specified dimensions.
func NewImage(width, height int) (img wandi.Image, err error) {
	panic("sfml.NewImage: not yet implemented.")
}

// LoadImage loads the provided image file and returns it as an image.
func LoadImage(imgPath string) (img wandi.Image, err error) {
	panic("sfml.LoadImage: not yet implemented.")
}

// ReadImage reads the provided image, converts it to the standard image format
// of this library and returns it.
func ReadImage(src image.Image) (img wandi.Image, err error) {
	panic("sfml.ReadImage: not yet implemented.")
}

// Draw draws the entire src image onto the dst image starting at the
// destination point dp.
func (dst *sfmlImg) Draw(dp image.Point, src wandi.Image) (err error) {
	panic("sfmlImg.Draw: not yet implemented.")
}

// DrawRect fills the destination rectangle dr of the dst image with
// corresponding pixels from the src image starting at the source point sp.
func (dst *sfmlImg) DrawRect(dr image.Rectangle, src wandi.Image, sp image.Point) (err error) {
	panic("sfmlImg.DrawRect: not yet implemented.")
}
