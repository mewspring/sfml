// Package texture handles hardware accelerated image drawing operations.
//
// This package uses a small subset of the features provided by the SFML library
// version 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package texture

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
	"github.com/mewmew/wandi/wandiutil"
)

// A Texture is a mutable collection of pixels, which is stored in GPU memory.
// It implements the wandi.Image interface.
type Texture struct {
	// A renderable GPU texture.
	RenderTex *C.sfRenderTexture
	// A SFML sprite representation of the GPU texture.
	Sprite *C.sfSprite
}

// New returns a new rendering texture of the specified dimensions. The texture
// is stored in GPU memory.
//
// Note: The Free method of the texture must be called when finished using it.
func New(width, height int) (img wandiutil.ImageClearFreer, err error) {
	return newTexture(width, height)
}

// newTexture returns a new rendering texture of the specified dimensions. The
// texture is stored in GPU memory.
//
// Note: The Free method of the texture must be called when finished using it.
func newTexture(width, height int) (tex *Texture, err error) {
	// Create a render texture of the specified dimensions.
	tex = new(Texture)
	tex.RenderTex = C.sfRenderTexture_create(C.uint(width), C.uint(height), C.sfFalse)
	if tex.RenderTex == nil {
		return nil, fmt.Errorf("texture.newTexture: unable to create %dx%d rendering texture", width, height)
	}

	// Create a sprite for the rendering texture.
	tex.Sprite = C.sfSprite_create()
	if tex.Sprite == nil {
		return nil, errors.New("texture.newTexture: unable to create sprite")
	}
	C.sfSprite_setTexture(tex.Sprite, tex.getNative(), C.sfTrue)

	return tex, nil
}

// getNative returns the GPU texture associated with the rendering texture.
func (tex *Texture) getNative() *C.sfTexture {
	return C.sfRenderTexture_getTexture(tex.RenderTex)
}

// Load loads the provided image file and returns it as a rendering texture. The
// texture is stored in GPU memory.
//
// Note: The Free method of the texture must be called when finished using it.
func Load(filePath string) (img wandiutil.ImageClearFreer, err error) {
	// Load source texture.
	texture := C.sfTexture_createFromFile(C.CString(filePath), nil)
	if texture == nil {
		return nil, fmt.Errorf("texture.Load: unable to load %q", filePath)
	}
	defer C.sfTexture_destroy(texture)

	// Create and return a rendering texture based on the source texture.
	return create(texture)
}

// create creates a new SFML rendering texture based on the provided SFML source
// texture.
func create(texture *C.sfTexture) (tex *Texture, err error) {
	// Create a sprite for the source texture.
	sprite := C.sfSprite_create()
	if sprite == nil {
		return nil, errors.New("texture.create: unable to create sprite")
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

// Create creates a new rendering texture based on the provided image and
// returns it. The texture is stored in GPU memory.
//
// Note: The Free method of the texture must be called when finished using it.
func Create(src image.Image) (img wandiutil.ImageClearFreer, err error) {
	switch srcImg := src.(type) {
	case *image.RGBA:
		width, height := srcImg.Rect.Dx(), srcImg.Rect.Dy()
		if srcImg.Stride != 4*width {
			// Use fallback conversion for subimages.
			return Create(convertImage(src))
		}
		// Create source texture based on the pixels of the image.
		texture := C.sfTexture_create(C.uint(width), C.uint(height))
		if texture == nil {
			return nil, fmt.Errorf("texture.Create: unable to create %dx%d source texture", width, height)
		}
		defer C.sfTexture_destroy(texture)
		pix := (*C.sfUint8)(unsafe.Pointer(&srcImg.Pix[0]))
		C.sfTexture_updateFromPixels(texture, pix, C.uint(width), C.uint(height), 0, 0)

		// Create and return a rendering texture based on the source texture.
		return create(texture)
	default:
		// Use fallback conversion for unknown image formats.
		log.Printf("texture.Create: using fallback for non-RGBA image format %T.\n", src)
		return Create(convertImage(src))
	}
}

// convertImage converts the provided src image to a RGBA image and returns it.
func convertImage(src image.Image) *image.RGBA {
	start := time.Now()
	bounds := src.Bounds()
	dr := image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	dst := image.NewRGBA(dr)
	draw.Draw(dst, dr, src, bounds.Min, draw.Src)
	log.Println("texture.convertImage: fallback conversion finished in:", time.Since(start))
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
		C.sfSprite_setTextureRect(srcImg.Sprite, sfmlIntRect(sr))
		C.sfSprite_setPosition(srcImg.Sprite, sfmlFloatPt(dp))
		C.sfRenderTexture_drawSprite(dst.RenderTex, srcImg.Sprite, nil)
		C.sfRenderTexture_display(dst.RenderTex)
	default:
		return fmt.Errorf("Texture.DrawRect: support for image format %T not yet implemented", src)
	}

	return nil
}
