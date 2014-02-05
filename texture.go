package sfml

// #include <SFML/Graphics/RenderTexture.h>
// #include <SFML/Graphics/Sprite.h>
import "C"

import (
	"errors"
	"image"

	"github.com/mewmew/wandi"
)

// A sfmlTexture is a mutable collection of pixels, which is stored in GPU
// memory. It implements the wandi.Image interface.
type sfmlTexture struct {
	tex    *C.sfRenderTexture
	sprite *C.sfSprite
}

// TODO(u): Provide more detailed error messages if possible.

// NewTexture returns a new image of the specified dimensions. The image texture
// is stored in GPU memory.
func NewTexture(width, height int) (img wandi.Image, err error) {
	sfmlTex := new(sfmlTexture)
	// TODO(u): Should we enable the depth buffer?
	sfmlTex.tex = C.sfRenderTexture_create(C.uint(width), C.uint(height), C.sfFalse)
	if sfmlTex.tex == nil {
		return nil, errors.New("sfml.NewTexture: unable to create texture")
	}
	sfmlTex.sprite = C.sfSprite_create()
	if sfmlTex.sprite == nil {
		return nil, errors.New("sfml.NewTexture: unable to create sprite")
	}
	C.sfSprite_setTexture(sfmlTex.sprite, sfmlTex.tex, C.sfTrue)
	return sfmlTex, nil
}

// LoadTexture loads the provided image file and returns it as an image. The
// image texture is stored in GPU memory.
func LoadTexture(imgPath string) (img wandi.Image, err error) {
	panic("sfml.LoadTexture: not yet implemented.")
}

// ReadTexture reads the provided image, converts it to the standard image
// format of this library and returns it. The image texture is stored in GPU
// memory.
func ReadTexture(src image.Image) (img wandi.Image, err error) {
	panic("sfml.ReadTexture: not yet implemented.")
}

// Free frees the image.
func (img *sfmlTexture) Free() {
	C.sfSprite_destroy(img.sprite)
	C.sfRenderTexture_destroy(img.tex)
}

// Draw draws the entire src image onto the dst image starting at the
// destination point dp.
func (dst *sfmlTexture) Draw(dp image.Point, src wandi.Image) (err error) {
	panic("sfmlTexture.Draw: not yet implemented.")
}

// DrawRect fills the destination rectangle dr of the dst image with
// corresponding pixels from the src image starting at the source point sp.
func (dst *sfmlTexture) DrawRect(dr image.Rectangle, src wandi.Image, sp image.Point) (err error) {
	panic("sfmlTexture.DrawRect: not yet implemented.")
}
