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
//
// sfMouseMoveEvent getMouseMoveEvent(sfEvent e) {
//    return e.mouseMove;
// }
//
// sfMouseButtonEvent getMouseButtonEvent(sfEvent e) {
//    return e.mouseButton;
// }
//
// sfMouseWheelEvent getMouseWheelEvent(sfEvent e) {
//    return e.mouseWheel;
// }
import "C"

import (
	"image"
	"log"

	"github.com/mewmew/we"
)

// PollEvent returns a pending event from the event queue or nil if the queue
// was empty. Note that more than one event may be present in the event queue.
//
// Note: PollEvent must be called from the same thread that created the window.
//
// Note: Some internal window events of SFML depend on calls to PollEvent to
// take effect. For instance a call to SetTitle will not update the window title
// until the next call of PollEvent.
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

// prev represents the previously recorded cursor position.
var prev image.Point

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
	case C.sfEvtMouseWheelMoved:
		e := C.getMouseWheelEvent(sfmlEvent)
		pt := image.Pt(int(e.x), int(e.y))
		event = we.ScrollY{
			Point: pt,
			Off:   int(e.delta),
			Mod:   getMod(),
		}
		return event
	case C.sfEvtMouseButtonPressed:
		e := C.getMouseButtonEvent(sfmlEvent)
		pt := image.Pt(int(e.x), int(e.y))
		event = we.MousePress{
			Point:  pt,
			Button: weButton(e.button),
			Mod:    getMod(),
		}
		return event
	case C.sfEvtMouseButtonReleased:
		e := C.getMouseButtonEvent(sfmlEvent)
		pt := image.Pt(int(e.x), int(e.y))
		event = we.MouseRelease{
			Point:  pt,
			Button: weButton(e.button),
			Mod:    getMod(),
		}
		return event
	case C.sfEvtMouseMoved:
		// TODO(u): implement MouseDrag event.
		e := C.getMouseMoveEvent(sfmlEvent)
		pt := image.Pt(int(e.x), int(e.y))
		event = we.MouseMove{
			Point: pt,
			From:  prev,
		}
		prev = pt
		return event
	case C.sfEvtMouseEntered:
		return we.MouseEnter(true)
	case C.sfEvtMouseLeft:
		return we.MouseEnter(false)

	default:
		log.Printf("weEvent: event %d not yet implemented.\n", typ)
	}

	// Ignore event.
	return nil
}
