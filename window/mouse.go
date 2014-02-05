package window

// #include <SFML/Window/Mouse.h>
import "C"

import (
	"github.com/mewmew/we"
)

// weButton returns the we.Button corresponding to the provided SFML mouse
// button.
func weButton(sfmlButton C.sfMouseButton) we.Button {
	switch sfmlButton {
	case C.sfMouseLeft:
		return we.ButtonLeft
	case C.sfMouseRight:
		return we.ButtonRight
	case C.sfMouseMiddle:
		return we.ButtonMiddle
	case C.sfMouseXButton1:
		return we.Button4
	case C.sfMouseXButton2:
		return we.Button5
	}
	return we.Button(sfmlButton)
}
