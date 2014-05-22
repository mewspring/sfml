package texture

// #cgo LDFLAGS: -lcsfml-graphics
// #include <SFML/Graphics.h>
import "C"

import (
	"image"
	"image/color"

	"github.com/mewmew/wandi"
)

// Drawable represent a drawable texture. It implements the wandi.Drawable and
// wandi.Image interfaces.
type Drawable struct {
	// A renderable GPU texture.
	tex *C.sfRenderTexture
	// A sprite representation of the GPU texture.
	sprite *C.sfSprite
}

// NewDrawable creates a drawable texture of the specified dimensions.
func NewDrawable(width, height int) (tex Drawable, err error) {
	panic("not yet implemented")
}

// texture returns the texture associated with the rendering texture.
func (tex *Drawable) texture() *C.sfTexture {
	return C.sfRenderTexture_getTexture(tex.tex)
}

// LoadDrawable loads the provided file and converts it into a drawable texture.
func LoadDrawable(path string) (tex Drawable, err error) {
	panic("not yet implemented")
}

// ReadDrawable reads the provided image and converts it into a drawable
// texture.
func ReadDrawable(img image.Image) (tex Drawable, err error) {
	panic("not yet implemented")
}

// Free frees the texture.
func (tex Drawable) Free() {
	panic("not yet implemented")
}

// Width returns the width of the texture.
func (tex Drawable) Width() int {
	panic("not yet implemented")
}

// Height returns the height of the texture.
func (tex Drawable) Height() int {
	panic("not yet implemented")
}

// Draw draws the entire src image onto the dst texture starting at the
// destination point dp.
func (dst Drawable) Draw(dp image.Point, src wandi.Image) (err error) {
	panic("not yet implemented")
}

// DrawRect draws a subset of the src image, as defined by the source rectangle
// sr, onto the dst texture starting at the destination point dp.
func (dst Drawable) DrawRect(dp image.Point, src wandi.Image, sr image.Rectangle) (err error) {
	panic("not yet implemented")
}

// Fill fills the entire texture with the provided color.
func (dst Drawable) Fill(c color.Color) {
	panic("not yet implemented")
}

// Image returns an image.Image representation of the texture.
func (tex Drawable) Image() (img image.Image, err error) {
	panic("not yet implemented")
}
