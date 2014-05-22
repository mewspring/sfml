package window

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
	sfRect := C.sfIntRect{
		left:   C.int(r.Min.X),
		top:    C.int(r.Min.Y),
		width:  C.int(r.Dx()),
		height: C.int(r.Dy()),
	}
	return sfRect
}

// sfmlFloatPt returns a SFML Vector2f based on the provided Go image.Point.
func sfmlFloatPt(pt image.Point) C.sfVector2f {
	sfPt := C.sfVector2f{
		x: C.float(pt.X),
		y: C.float(pt.Y),
	}
	return sfPt
}

// utf32 returns the UTF-32 representation of s.
func utf32(s string) *C.sfUint32 {
	s32 := make([]C.sfUint32, 0, len(s))
	for _, r := range s {
		s32 = append(s32, C.sfUint32(r))
	}
	// End with a NULL byte.
	s32 = append(s32, 0)
	return &s32[0]
}
