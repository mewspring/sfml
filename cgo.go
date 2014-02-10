package sfml

// #cgo LDFLAGS: -lcsfml-graphics
// #include <SFML/Graphics.h>
import "C"

import (
	"image"
	"image/color"
)

// sfmlColor returns a SFML Color based on the provided Go color.Color.
func sfmlColor(c color.Color) C.sfColor {
	r, g, b, a := c.RGBA()
	sfColor := C.sfColor{
		r: C.sfUint8(r),
		g: C.sfUint8(g),
		b: C.sfUint8(b),
		a: C.sfUint8(a),
	}
	return sfColor
}

// sfmlIntRect returns a SFML IntRect based on the provided Go image.Rectangle.
func sfmlIntRect(r image.Rectangle) C.sfIntRect {
	rect := C.sfIntRect{
		left:   C.int(r.Min.X),
		top:    C.int(r.Min.Y),
		width:  C.int(r.Dx()),
		height: C.int(r.Dy()),
	}
	return rect
}

// sfmlFloatPt returns a SFML Vector2f based on the provided Go image.Point.
func sfmlFloatPt(pt image.Point) C.sfVector2f {
	sfPt := C.sfVector2f{
		x: C.float(pt.X),
		y: C.float(pt.Y),
	}
	return sfPt
}
