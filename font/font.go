// Package font handles text rendering based on the size, style and color of
// fonts.
//
// This package uses a small subset of the features provided by the SFML library
// version 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package font

// #include <SFML/Graphics.h>
import "C"

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/mewmew/sfml/texture"
)

// A Font describes a particular size, style and color in which to render text.
type Font struct {
	// font provides the glyphs (visual characters) of the font.
	font *C.sfFont
	// text specifies the size, style and color of the rendered text.
	text *C.sfText
}

// Load loads the provided TTF font. The default font size, style and color is
// 12, regular (no style) and black respectively.
//
// Note: The Free method of the font must be called when finished using it.
func Load(filePath string) (f *Font, err error) {
	// Load the FFT font file.
	f = new(Font)
	f.font = C.sfFont_createFromFile(C.CString(filePath))
	if f.font == nil {
		return nil, fmt.Errorf("font.Load: unable to load %q", filePath)
	}

	// Create a SFML text and associate it with the font.
	f.text = C.sfText_create()
	if f.text == nil {
		return nil, errors.New("font.Load: unable to create text")
	}
	C.sfText_setFont(f.text, f.font)

	// Set the default size, style and color of the font.
	f.SetSize(12)
	f.SetStyle(Regular)
	f.SetColor(color.Black)

	return f, nil
}

// Free frees the font.
func (f *Font) Free() {
	C.sfText_destroy(f.text)
	C.sfFont_destroy(f.font)
}

// SetSize sets the size in pixels of the text drawn with font.
func (f *Font) SetSize(size int) {
	C.sfText_setCharacterSize(f.text, C.uint(size))
}

// Style is a bitfield which represents the style of a font.
type Style uint32

// Font styles.
const (
	// Regular characters (no style).
	Regular Style = C.sfTextRegular
	// Bold characters.
	Bold Style = C.sfTextBold
	// Italic characters.
	Italic Style = C.sfTextItalic
	// Underlined characters.
	Underlined Style = C.sfTextUnderlined
)

// SetStyle sets the style of the text drawn with font.
func (f *Font) SetStyle(style Style) {
	C.sfText_setStyle(f.text, C.sfUint32(style))
}

// SetColor sets the color of the text drawn with font.
func (f *Font) SetColor(c color.Color) {
	C.sfText_setColor(f.text, sfmlColor(c))
}

// Render renders an image of the provided text in the style, size and color of
// the font.
//
// Note: The Free method of the returned image must be called when finished
// using it.
func (f *Font) Render(s string) (img texture.Image, err error) {
	// Specify the text to be rendered.
	C.sfText_setUnicodeString(f.text, utf32(s))

	// Ignore padding.
	bounds := C.sfText_getGlobalBounds(f.text)
	pt := C.sfVector2f{
		x: -bounds.left,
		y: -bounds.top,
	}
	C.sfText_setPosition(f.text, pt)

	// Create a rendering texture of the same dimensions as the text.
	tex, err := texture.New(int(bounds.width), int(bounds.height))
	if err != nil {
		return nil, err
	}

	// Draw the text onto the rendering texture.
	dst := tex.(*texture.Texture)
	C.sfRenderTexture_drawText(dst.RenderTex, f.text, nil)
	C.sfRenderTexture_display(dst.RenderTex)

	return dst, nil
}
