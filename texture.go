package sfml

// #include <SFML/Graphics.h>
import "C"

import (
	"errors"
	"image"

	"github.com/mewmew/wandi"
)

// A Texture is a mutable collection of pixels, which is stored in GPU memory.
// It implements the wandi.Image interface.
type Texture struct {
	// A renderable GPU texture.
	RenderTex *C.sfRenderTexture
	// A SFML sprite representation of the GPU texture.
	Sprite *C.sfSprite
}

// TODO(u): Provide more detailed error messages if possible.

// NewTexture returns a new image of the specified dimensions. The texture is
// stored in GPU memory.
//
// Note: The Free method of the image should be called when finished using it.
func NewTexture(width, height int) (img wandi.Image, err error) {
	tex := new(Texture)
	// TODO(u): Should we enable the depth buffer?
	tex.RenderTex = C.sfRenderTexture_create(C.uint(width), C.uint(height), C.sfFalse)
	if tex.RenderTex == nil {
		return nil, errors.New("sfml.NewTexture: unable to create texture")
	}
	tex.Sprite = C.sfSprite_create()
	if tex.Sprite == nil {
		return nil, errors.New("sfml.NewTexture: unable to create sprite")
	}
	t := C.sfRenderTexture_getTexture(tex.RenderTex)
	C.sfSprite_setTexture(tex.Sprite, t, C.sfTrue)
	return tex, nil
}

// LoadTexture loads the provided image file and returns it as an image. The
// texture is stored in GPU memory.
//
// Note: The Free method of the image should be called when finished using it.
func LoadTexture(filePath string) (img wandi.Image, err error) {
	panic("sfml.LoadTexture: not yet implemented.")
}

// ReadTexture reads the provided image, converts it to the standard image
// format of this library and returns it. The texture is stored in GPU memory.
//
// Note: The Free method of the image should be called when finished using it.
func ReadTexture(src image.Image) (img wandi.Image, err error) {
	panic("sfml.ReadTexture: not yet implemented.")
}

// Free frees the image.
func (tex *Texture) Free() {
	C.sfSprite_destroy(tex.Sprite)
	C.sfRenderTexture_destroy(tex.RenderTex)
}

// Width returns the width of the image.
func (tex *Texture) Width() int {
	size := C.sfRenderTexture_getSize(tex.RenderTex)
	return int(size.x)
}

// Height returns the height of the image.
func (tex *Texture) Height() int {
	size := C.sfRenderTexture_getSize(tex.RenderTex)
	return int(size.y)
}

// Draw draws the entire src image onto the dst image starting at the
// destination point dp.
func (dst *Texture) Draw(dp image.Point, src wandi.Image) (err error) {
	panic("Texture.Draw: not yet implemented.")
}

// TODO(u): Rethink dr, sp for DrawRect. It would probably be more consistent to
// use dp, sr instead.

// DrawRect fills the destination rectangle dr of the dst image with
// corresponding pixels from the src image starting at the source point sp.
func (dst *Texture) DrawRect(dr image.Rectangle, src wandi.Image, sp image.Point) (err error) {
	switch srcImg := src.(type) {
	case *Texture:
		sr := image.Rect(sp.X, sp.Y, sp.X+dr.Dx(), sp.Y+dr.Dy())
		C.sfSprite_setTextureRect(srcImg.Sprite, sfmlIntRect(sr))
		dp := dr.Min
		C.sfSprite_setPosition(srcImg.Sprite, sfmlFloatPt(dp))
		C.sfRenderTexture_drawSprite(dst.RenderTex, srcImg.Sprite, nil)
		C.sfRenderTexture_display(dst.RenderTex)
	}

	return nil
}
