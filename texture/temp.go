// TODO(u): This file is intended to facilitate development. Remove it when no
// longer needed.

package texture

// #include <SFML/Graphics.h>
import "C"

import (
	"errors"
	"fmt"
)

// WriteFile writes the texture to an image file.
//
// TODO(u): This method is intended to facilitate development. Remove it when no
// longer needed.
func (tex *Texture) WriteFile(filePath string) (err error) {
	img := C.sfTexture_copyToImage(tex.getNative())
	if img == nil {
		return errors.New("Texture.WriteFile: unable to create image of texture")
	}
	if C.sfImage_saveToFile(img, C.CString(filePath)) == C.sfFalse {
		return fmt.Errorf("Texture.WriteFile: unable to write image to %q", filePath)
	}
	return nil
}
