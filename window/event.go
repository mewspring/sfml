package window

// #include <SFML/Graphics/RenderWindow.h>
//
// sfEventType getEventType(sfEvent e) {
//    return e.type;
// }
//
// sfSizeEvent getSizeEvent(sfEvent e) {
//    return e.size;
// }
//
// sfKeyEvent getKeyEvent(sfEvent e) {
//    return e.key;
// }
//
// sfTextEvent getTextEvent(sfEvent e) {
//    return e.text;
// }
import "C"

import (
	"log"

	"github.com/mewmew/we"
)

// TODO(u): Add note about the need to call PollEvent even when not using
// events, since this is required by the window.

// PollEvent returns a pending event from the event queue or nil if the queue
// was empty.
//
// Note: PollEvent must be called from the same thread that created the window.
func (sfmlWin *sfmlWindow) PollEvent() (event we.Event) {
	// Poll the event queue until we locate a non-nil event or the queue is
	// empty.
	var sfmlEvent C.sfEvent
	for {
		if C.sfRenderWindow_pollEvent(sfmlWin.w, &sfmlEvent) == C.sfFalse {
			// Return nil if the event queue is empty.
			return nil
		}
		event = weEvent(sfmlEvent)
		if event != nil {
			return event
		}
	}
}

// weEvent returns the we.Event corresponding to the provided SFML event.
func weEvent(sfmlEvent C.sfEvent) (event we.Event) {
	typ := C.getEventType(sfmlEvent)
	switch typ {
	// Window events.
	case C.sfEvtClosed:
		return we.Close{}
	case C.sfEvtResized:
		e := C.getSizeEvent(sfmlEvent)
		event = we.Resize{
			Width:  int(e.width),
			Height: int(e.height),
		}
		return event

	// Keyboard events.
	case C.sfEvtTextEntered:
		e := C.getTextEvent(sfmlEvent)
		event = we.KeyRune(e.unicode)
		return event
	case C.sfEvtKeyPressed:
		e := C.getKeyEvent(sfmlEvent)
		event = we.KeyPress{
			Key: weKey(e.code),
			Mod: weMod(e),
		}
		return event
	case C.sfEvtKeyReleased:
		e := C.getKeyEvent(sfmlEvent)
		event = we.KeyRelease{
			Key: weKey(e.code),
			Mod: weMod(e),
		}
		return event

	// Mouse events.
	//case C.sfEvtMouseWheelMoved:
	//case C.sfEvtMouseButtonPressed:
	//case C.sfEvtMouseButtonReleased:
	//case C.sfEvtMouseMoved:
	//case C.sfEvtMouseEntered:
	//case C.sfEvtMouseLeft:

	default:
		log.Printf("weEvent: event %d not yet implemented.\n", typ)
	}

	// Ignore event.
	return nil
}
