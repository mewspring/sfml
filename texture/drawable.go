package texture

// #cgo LDFLAGS: -lcsfml-graphics
// #include <SFML/Graphics.h>
import "C"

import (
	"errors"
	"fmt"
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
//
// Note: The Free method of the texture must be called when finished using it.
func NewDrawable(width, height int) (tex Drawable, err error) {
	// Create a rendering texture of the specified dimensions.
	tex.tex = C.sfRenderTexture_create(C.uint(width), C.uint(height), C.sfFalse)
	if tex.tex == nil {
		return Drawable{}, fmt.Errorf("texture.NewDrawable: unable to create %dx%d rendering texture", width, height)
	}

	// Create a sprite for the rendering texture.
	tex.sprite = C.sfSprite_create()
	if tex.sprite == nil {
		return Drawable{}, errors.New("texture.NewDrawable: unable to create sprite")
	}
	C.sfSprite_setTexture(tex.sprite, tex.texture(), C.sfTrue)

	return tex, nil
}

// texture returns the texture associated with the rendering texture.
func (tex *Drawable) texture() *C.sfTexture {
	return C.sfRenderTexture_getTexture(tex.tex)
}

// LoadDrawable loads the provided file and converts it into a drawable texture.
//
// Note: The Free method of the texture must be called when finished using it.
func LoadDrawable(path string) (tex Drawable, err error) {
	// Load the read-only source texture from file.
	src, err := Load(path)
	if err != nil {
		return Drawable{}, err
	}
	defer src.Free()

	// Create a drawable texture of the same image dimensions.
	tex, err = NewDrawable(src.Width(), src.Height())
	if err != nil {
		return Drawable{}, err
	}

	// Draw the read-only source texture onto the drawable texture.
	err = tex.Draw(image.ZP, src)
	if err != nil {
		return Drawable{}, err
	}

	return tex, nil
}

// ReadDrawable reads the provided image and converts it into a drawable
// texture.
//
// Note: The Free method of the texture must be called when finished using it.
func ReadDrawable(img image.Image) (tex Drawable, err error) {
	panic("not yet implemented")
}

// Free frees the texture.
func (tex Drawable) Free() {
	C.sfSprite_destroy(tex.sprite)
	C.sfRenderTexture_destroy(tex.tex)
}

// Width returns the width of the texture.
func (tex Drawable) Width() int {
	size := C.sfRenderTexture_getSize(tex.tex)
	return int(size.x)
}

// Height returns the height of the texture.
func (tex Drawable) Height() int {
	size := C.sfRenderTexture_getSize(tex.tex)
	return int(size.y)
}

// Draw draws the entire src image onto the dst texture starting at the
// destination point dp.
func (dst Drawable) Draw(dp image.Point, src wandi.Image) (err error) {
	sr := image.Rect(0, 0, src.Width(), src.Height())
	return dst.DrawRect(dp, src, sr)
}

// DrawRect draws a subset of the src image, as defined by the source rectangle
// sr, onto the dst texture starting at the destination point dp.
func (dst Drawable) DrawRect(dp image.Point, src wandi.Image, sr image.Rectangle) (err error) {
	switch srcImg := src.(type) {
	case *Drawable:
		C.sfSprite_setTextureRect(srcImg.sprite, sfmlIntRect(sr))
		C.sfSprite_setPosition(srcImg.sprite, sfmlFloatPt(dp))
		C.sfRenderTexture_drawSprite(dst.tex, srcImg.sprite, nil)
		C.sfRenderTexture_display(dst.tex)
	case *Image:
		C.sfSprite_setTextureRect(srcImg.sprite, sfmlIntRect(sr))
		C.sfSprite_setPosition(srcImg.sprite, sfmlFloatPt(dp))
		C.sfRenderTexture_drawSprite(dst.tex, srcImg.sprite, nil)
		C.sfRenderTexture_display(dst.tex)
	default:
		return fmt.Errorf("Texture.DrawRect: support for image format %T not yet implemented", src)
	}
	return nil
}

// Fill fills the entire texture with the provided color.
func (tex Drawable) Fill(c color.Color) {
	C.sfRenderTexture_clear(tex.tex, sfmlColor(c))
}

// Image returns an image.Image representation of the texture.
func (tex Drawable) Image() (img image.Image, err error) {
	panic("not yet implemented")
}
