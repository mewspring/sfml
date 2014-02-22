// Package font handles graphical text entries with customizable font size,
// style and color.
//
// This package uses a small subset of the features provided by the SFML library
// version 2.1 [1].
//
// [1]: http://www.sfml-dev.org/
package font

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
func Load(filePath string) (f *Font, err error) {
	// Load the FFT font file.
	f = new(Font)
	f.font = C.sfFont_createFromFile(C.CString(filePath))
	if f.font == nil {
		return nil, fmt.Errorf("font.Load: unable to load %q", filePath)
	}

	return f, nil
}

// Free frees the font.
func (f *Font) Free() {
	C.sfFont_destroy(f.font)
}
