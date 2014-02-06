// TODO(u): This file is intended to facilitate development. Remove it when no
// longer needed.

package sfml

// #include <SFML/Graphics.h>
import "C"

import (
	"errors"
	"fmt"
)

func (img *Image) WriteFile(filePath string) (err error) {
	if C.sfImage_saveToFile(img.Img, C.CString(filePath)) == C.sfFalse {
		return fmt.Errorf("Image.WriteFile: unable to write image to %q", filePath)
	}
	return nil
}

func (tex *Texture) WriteFile(filePath string) (err error) {
	t := C.sfRenderTexture_getTexture(tex.Tex)
	img := C.sfTexture_copyToImage(t)
	if img == nil {
		return errors.New("Texture.WriteFile: unable to create image of texture.")
	}
	if C.sfImage_saveToFile(img, C.CString(filePath)) == C.sfFalse {
		return fmt.Errorf("Image.WriteFile: unable to write image to %q", filePath)
	}
	return nil
}
