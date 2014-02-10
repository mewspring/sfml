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
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"
	"unsafe"

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

// NewTexture returns a new rendering texture of the specified dimensions. The
// texture is stored in GPU memory.
//
// Note: The Free method of the texture should be called when finished using it.
func NewTexture(width, height int) (img wandi.Image, err error) {
	return newTexture(width, height)
}

// newTexture returns a new rendering texture of the specified dimensions. The
// texture is stored in GPU memory.
//
// Note: The Free method of the texture should be called when finished using it.
func newTexture(width, height int) (tex *Texture, err error) {
	// Create a render texture of the specified dimensions.
	tex = new(Texture)
	tex.RenderTex = C.sfRenderTexture_create(C.uint(width), C.uint(height), C.sfFalse)
	if tex.RenderTex == nil {
		return nil, errors.New("sfml.newTexture: unable to create texture")
	}

	// Create a sprite for the rendering texture.
	tex.Sprite = C.sfSprite_create()
	if tex.Sprite == nil {
		return nil, errors.New("sfml.newTexture: unable to create sprite")
	}
	C.sfSprite_setTexture(tex.Sprite, tex.getNative(), C.sfTrue)

	return tex, nil
}

// getNative returns the GPU texture associated with the rendering texture.
func (tex *Texture) getNative() *C.sfTexture {
	return C.sfRenderTexture_getTexture(tex.RenderTex)
}

// LoadTexture loads the provided image file and returns it as a rendering
// texture. The texture is stored in GPU memory.
//
// Note: The Free method of the texture should be called when finished using it.
func LoadTexture(filePath string) (img wandi.Image, err error) {
	// Load source texture.
	texture := C.sfTexture_createFromFile(C.CString(filePath), nil)
	if texture == nil {
		return nil, fmt.Errorf("sfml.LoadTexture: unable to load %q", filePath)
	}
	defer C.sfTexture_destroy(texture)

	return createRenderTexture(texture)
}

// createRenderTexture creates a new SFML rendering texture based on the
// provided SFML source texture.
func createRenderTexture(texture *C.sfTexture) (tex *Texture, err error) {
	// Create sprite for the source texture.
	sprite := C.sfSprite_create()
	if sprite == nil {
		return nil, errors.New("sfml.createRenderTexture: unable to create sprite")
	}
	defer C.sfSprite_destroy(sprite)
	C.sfSprite_setTexture(sprite, texture, C.sfTrue)

	// Create the rendering texture.
	size := C.sfTexture_getSize(texture)
	width, height := int(size.x), int(size.y)
	tex, err = newTexture(width, height)
	if err != nil {
		return nil, err
	}

	// Draw the source sprite onto the rendering texture.
	C.sfRenderTexture_drawSprite(tex.RenderTex, sprite, nil)
	C.sfRenderTexture_display(tex.RenderTex)

	return tex, nil
}

// ReadTexture reads the provided image, converts it to a rendering texture and
// returns it. The texture is stored in GPU memory.
//
// Note: The Free method of the texture should be called when finished using it.
func ReadTexture(src image.Image) (img wandi.Image, err error) {
	switch srcImg := src.(type) {
	case *image.RGBA:
		width, height := srcImg.Rect.Dx(), srcImg.Rect.Dy()
		if srcImg.Stride != 4*width {
			// Use fallback conversion for subimages.
			return ReadTexture(convertImage(src))
		}
		// Create source texture based on the pixels of the image.
		texture := C.sfTexture_create(C.uint(width), C.uint(height))
		defer C.sfTexture_destroy(texture)
		pix := (*C.sfUint8)(unsafe.Pointer(&srcImg.Pix[0]))
		C.sfTexture_updateFromPixels(texture, pix, C.uint(width), C.uint(height), 0, 0)

		return createRenderTexture(texture)
	default:
		log.Printf("sfml.ReadTexture: using fallback for non-RGBA image format %T.\n", src)
		// Use fallback conversion for unknown image formats.
		return ReadTexture(convertImage(src))
	}
}

// convertImage converts the provided src image to a RGBA image and returns it.
func convertImage(src image.Image) (dst *image.RGBA) {
	start := time.Now()
	bounds := src.Bounds()
	dr := image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	dst = image.NewRGBA(dr)
	draw.Draw(dst, dr, src, bounds.Min, draw.Src)
	log.Println("convertImage: fallback conversion finished in:", time.Since(start))
	return dst
}

// Free frees the texture.
func (tex *Texture) Free() {
	C.sfSprite_destroy(tex.Sprite)
	C.sfRenderTexture_destroy(tex.RenderTex)
}

// Width returns the width of the texture.
func (tex *Texture) Width() int {
	size := C.sfRenderTexture_getSize(tex.RenderTex)
	return int(size.x)
}

// Height returns the height of the texture.
func (tex *Texture) Height() int {
	size := C.sfRenderTexture_getSize(tex.RenderTex)
	return int(size.y)
}

// Clear clears the texture and fills it with the provided color.
func (tex *Texture) Clear(c color.Color) {
	C.sfRenderTexture_clear(tex.RenderTex, sfmlColor(c))
}

// Draw draws the entire src image onto the dst image starting at the
// destination point dp.
func (dst *Texture) Draw(dp image.Point, src wandi.Image) (err error) {
	sr := image.Rect(0, 0, src.Width(), src.Height())
	return dst.DrawRect(dp, src, sr)
}

// DrawRect fills the destination rectangle dr of the dst image with
// corresponding pixels from the src image starting at the source point sp.
func (dst *Texture) DrawRect(dp image.Point, src wandi.Image, sr image.Rectangle) (err error) {
	switch srcImg := src.(type) {
	case *Texture:
		drawRectSprite(dst.RenderTex, dp, srcImg.Sprite, sr)
	default:
		return fmt.Errorf("Texture.DrawRect: support for image format %T not yet implemented", src)
	}

	return nil
}

// drawRectSprite fills the destination rectangle dr of the dst SFML render
// texture with corresponding pixels from the src SFML sprite starting at the
// source point sp.
func drawRectSprite(dst *C.sfRenderTexture, dp image.Point, src *C.sfSprite, sr image.Rectangle) {
	C.sfSprite_setTextureRect(src, sfmlIntRect(sr))
	C.sfSprite_setPosition(src, sfmlFloatPt(dp))
	C.sfRenderTexture_drawSprite(dst, src, nil)
	C.sfRenderTexture_display(dst)
}
