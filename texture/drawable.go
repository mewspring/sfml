package texture

// #cgo LDFLAGS: -lcsfml-graphics
// #include <string.h>
// #include <SFML/Graphics.h>
import "C"

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"unsafe"

	"github.com/mewmew/sfml/font"
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
func (tex Drawable) texture() *C.sfTexture {
	return C.sfRenderTexture_getTexture(tex.tex)
}

// LoadDrawable loads the provided file and converts it into a drawable texture.
//
// Note: The Free method of the texture must be called when finished using it.
func LoadDrawable(path string) (tex Drawable, err error) {
	// Load the provided file and convert it into a read-only texture.
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
	// Read the provided image and convert it into a read-only texture.
	src, err := Read(img)
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
	case Drawable:
		C.sfSprite_setTextureRect(srcImg.sprite, sfmlIntRect(sr))
		C.sfSprite_setPosition(srcImg.sprite, sfmlFloatPt(dp))
		C.sfRenderTexture_drawSprite(dst.tex, srcImg.sprite, nil)
		C.sfRenderTexture_display(dst.tex)
	case Image:
		C.sfSprite_setTextureRect(srcImg.sprite, sfmlIntRect(sr))
		C.sfSprite_setPosition(srcImg.sprite, sfmlFloatPt(dp))
		C.sfRenderTexture_drawSprite(dst.tex, srcImg.sprite, nil)
		C.sfRenderTexture_display(dst.tex)
	case font.Text:
		// TODO(u): Handle sr?
		text := textText(srcImg)
		C.sfText_setPosition(text, sfmlFloatPt(dp))
		C.sfRenderTexture_drawText(dst.tex, text, nil)
		C.sfRenderTexture_display(dst.tex)
	default:
		return fmt.Errorf("Drawable.DrawRect: support for image format %T not yet implemented", src)
	}
	return nil
}

// Fill fills the entire texture with the provided color.
func (tex Drawable) Fill(c color.Color) {
	C.sfRenderTexture_clear(tex.tex, sfmlColor(c))
}

// Image returns an image.Image representation of the texture.
func (tex Drawable) Image() (img image.Image, err error) {
	// Copy the rendering texture to a SFML image.
	sfImg := C.sfTexture_copyToImage(tex.texture())
	if sfImg == nil {
		return nil, errors.New("Drawable.Image: unable to create image from texture")
	}
	defer C.sfImage_destroy(sfImg)

	// Create a Go RGBA image based on the pixels of the SFML image.
	pix := C.sfImage_getPixelsPtr(sfImg)
	if pix == nil {
		return nil, errors.New("Drawable.Image: unable to locate image pixels")
	}
	size := C.sfImage_getSize(sfImg)
	dst := image.NewRGBA(image.Rect(0, 0, int(size.x), int(size.y)))
	C.memcpy(unsafe.Pointer(&dst.Pix[0]), unsafe.Pointer(pix), C.size_t(len(dst.Pix)))

	return dst, nil
}
