package font

// #include <SFML/Graphics.h>
import "C"

import (
	"errors"
	"image/color"
)

// A Text is a graphical text entry with a specific font size, style and color.
// It implements the wandi.Image interface.
type Text struct {
	// A graphical text entry.
	Text *C.sfText
}

// NewText returns a new graphical text entry based on the provided font and any
// optional customization arguments. The initial text, size, style and color of
// the graphical text entry can be customized through string, int, Style and
// color.Color arguments respectively, depending on the type of the argument.
//
// The default font size, style and color of the text is 12, regular (no style)
// and black respectively.
//
// Note: The Free method of the graphical text entry must be called when
// finished using it.
func NewText(f *Font, args ...interface{}) (text *Text, err error) {
	// Create a SFML text and associate the font with it.
	text = new(Text)
	text.Text = C.sfText_create()
	if text.Text == nil {
		return nil, errors.New("font.NewText: unable to create text")
	}
	C.sfText_setFont(text.Text, f.font)

	// Set the default font size, style and color of the text.
	text.SetSize(12)
	text.SetStyle(Regular)
	text.SetColor(color.Black)

	// Customize the text, size, style and color based on the provided arguments.
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			text.SetText(v)
		case int:
			text.SetSize(v)
		case Style:
			text.SetStyle(v)
		case color.Color:
			text.SetColor(v)
		}
	}

	return text, nil
}

// Free frees the graphical text entry.
func (text *Text) Free() {
	C.sfText_destroy(text.Text)
}

// SetText sets the text of the graphical text entry.
func (text *Text) SetText(s string) {
	C.sfText_setUnicodeString(text.Text, utf32(s))
}

// SetSize sets the font size of the text, in pixels.
func (text *Text) SetSize(size int) {
	C.sfText_setCharacterSize(text.Text, C.uint(size))
}

// Style is a bitfield which represents the style of a text.
type Style uint32

// Text styles.
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

// SetStyle sets the style of the text.
func (text *Text) SetStyle(style Style) {
	C.sfText_setStyle(text.Text, C.sfUint32(style))
}

// SetColor sets the color of the text.
func (text *Text) SetColor(c color.Color) {
	C.sfText_setColor(text.Text, sfmlColor(c))
}

// Width returns the width of the graphical text entry.
func (text *Text) Width() int {
	bounds := C.sfText_getLocalBounds(text.Text)
	return int(bounds.width + bounds.left)
}

// Height returns the height of the graphical text entry.
func (text *Text) Height() int {
	bounds := C.sfText_getLocalBounds(text.Text)
	return int(bounds.height + bounds.top)
}
