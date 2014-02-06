package sfml

// #cgo LDFLAGS: -lcsfml-graphics
// #include <SFML/Graphics.h>
import "C"

import (
	"image"
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

// sfmlIntRect returns a C sfIntRect based on the provided Go image.Rectangle.
func sfmlIntRect(r image.Rectangle) C.sfIntRect {
	rect := C.sfIntRect{
		left:   C.int(r.Min.X),
		top:    C.int(r.Min.Y),
		width:  C.int(r.Dx()),
		height: C.int(r.Dy()),
	}
	return rect
}

// sfmlFloatRect returns a C sfFloatRect based on the provided Go
// image.Rectangle.
func sfmlFloatRect(r image.Rectangle) C.sfFloatRect {
	sfmlRect := C.sfFloatRect{
		left:   C.float(r.Min.X),
		top:    C.float(r.Min.Y),
		width:  C.float(r.Dx()),
		height: C.float(r.Dy()),
	}
	return sfmlRect
}

// sfmlView defines a camera in the 2D scene. A view is composed of a source
// rectangle, which defines what part of the 2D scene is shown, and a target
// viewport, which defines where the contents of the source rectangle will be
// displayed on the render target.
type sfmlView struct {
	view *C.sfView
}

// newView returns a C sfView based on the provided destination rectangle and
// source point. The scale between the view and the viewport is one to one.
//
// Note: The free method of the view must be called when finished using it.
func newView(dr image.Rectangle, sp image.Point) (view *sfmlView) {
	view = new(sfmlView)
	sr := image.Rect(sp.X, sp.Y, sp.X+dr.Dx(), sp.Y+dr.Dy())
	srcRect := sfmlFloatRect(sr)
	view.view = C.sfView_createFromRect(srcRect)
	viewportRect := sfmlFloatRect(dr)
	C.sfView_setViewport(view.view, viewportRect)
	return view
}

// free frees the view.
func (view *sfmlView) free() {
	C.sfView_destroy(view.view)
}
