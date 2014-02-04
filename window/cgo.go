package window

// #include <SFML/Graphics/Color.h>
import "C"

import (
	"image/color"
)

// sfmlColor returns a C sfColor based on the provided Go color.Color.
func sfmlColor(c color.Color) C.sfColor {
	r, g, b, a := c.RGBA()
	sfmlCol := C.sfColor{
		r: C.sfUint8(r),
		g: C.sfUint8(g),
		b: C.sfUint8(b),
		a: C.sfUint8(a),
	}
	return sfmlCol
}
