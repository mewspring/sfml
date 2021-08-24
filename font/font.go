// Package font handles graphical text entries with customizable font size,
// style and color. It uses a small subset of the features provided by the SFML
// library version 2.5 [1].
//
// [1]: http://www.sfml-dev.org/
package font

// #cgo LDFLAGS: -lcsfml-graphics
// #include <SFML/Graphics.h>
import "C"

import (
	"fmt"
)

// A Font provides glyphs (visual characters) and metrics used for text
// rendering.
type Font struct {
	// A TTF font.
	font *C.sfFont
}

// Load loads the provided TTF font.
//
// Note: The Free method of the font must be called when finished using it.
func Load(path string) (f Font, err error) {
	// Load the FFT font file.
	f.font = C.sfFont_createFromFile(C.CString(path))
	if f.font == nil {
		return Font{}, fmt.Errorf("font.Load: unable to load %q", path)
	}

	return f, nil
}

// Free frees the font.
func (f Font) Free() {
	C.sfFont_destroy(f.font)
}
