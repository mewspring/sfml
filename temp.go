// TODO(u): This file is intended to facilitate development. Remove it when no
// longer needed.

package sfml

// #include <SFML/Graphics.h>
import "C"

import (
	"errors"
	"fmt"
	"image/color"
)

// WriteFile writes the texture to an image file.
func (tex *Texture) WriteFile(filePath string) (err error) {
	img := C.sfTexture_copyToImage(tex.getTex())
	if img == nil {
		return errors.New("Texture.WriteFile: unable to create image of texture")
	}
	if C.sfImage_saveToFile(img, C.CString(filePath)) == C.sfFalse {
		return fmt.Errorf("Texture.WriteFile: unable to write image to %q", filePath)
	}
	return nil
}

// Clear clears the texture and fills it with the provided color.
func (tex *Texture) Clear(c color.Color) {
	C.sfRenderTexture_clear(tex.RenderTex, sfmlColor(c))
}
