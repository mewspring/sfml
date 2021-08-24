package window

// #include <SFML/Window.h>
import "C"

import (
	"log"

	"github.com/mewspring/we"
)

// weButton returns the we.Button corresponding to the provided SFML mouse
// button.
func weButton(button C.sfMouseButton) we.Button {
	switch button {
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

	// Unknown mouse button.
	log.Printf("window.weButton: unknown mouse button %d.\n", button)
	return 0
}
