package window

// #include <SFML/Graphics/RenderWindow.h>
import "C"

import (
	"github.com/mewmew/we"
)

// TODO(u): Add note about the need to call PollEvent even when not using
// events, since this is required by the window.

// PollEvent returns a pending event from the event queue or nil if the queue
// was empty.
//
// Note: PollEvent must be called from the same thread that created the window.
func (sfmlWin *sfmlWindow) PollEvent() (event we.Event) {
	var sfmlEvent C.sfEvent
	if C.sfRenderWindow_pollEvent(sfmlWin.w, &sfmlEvent) == C.sfFalse {
		return nil
	}
	// TODO(u): not yet implemented.
	return nil
}
