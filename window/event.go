package win

import (
	"github.com/mewmew/we"
)

// TODO(u): Add note about the need to call PollEvent even when not using
// events, since this is required by the window.

// PollEvent returns a pending event from the event queue or nil if the queue
// was empty.
//
// Note: PollEvent must be called from the same thread that created the window.
func (win *sfmlWindow) PollEvent() (event we.Event) {
	panic("sfmlWindow.PollEvent: not yet implemented.")
}
