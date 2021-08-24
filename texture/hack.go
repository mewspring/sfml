package texture

// #include <SFML/Graphics.h>
import "C"

import (
	"unsafe"

	"github.com/mewspring/sfml/font"
)

// textHack is a copy of font.Text without modifications. Through the use of
// unsafe and with knowledge of its memory layout we are able to access
// unexported members. This hack allows us to cross package barriers while
// keeping the exported API clean.
type textHack struct {
	// A graphical text entry.
	text *C.sfText
}

// textText returns the text of the provided font.Text.
func textText(text font.Text) *C.sfText {
	return (*textHack)(unsafe.Pointer(&text)).text
}
