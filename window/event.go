package window

// #include <SFML/Graphics.h>
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
// Note: The main thread must be used for both window creation and event
// handling. It is perfectly fine to use separate threads for rendering and
// event handling, as long as all event handling takes place in the main thread.
//
// Note: Some internal window events of SFML depend on calls to PollEvent to
// take effect. For instance a call to SetTitle will not update the window title
// until the next call of PollEvent.
func (win *Window) PollEvent() (event we.Event) {
	// Poll the event queue until we locate a non-nil event or the queue is
	// empty.
	var sfEvent C.sfEvent
	for {
		if C.sfRenderWindow_pollEvent(win.Win, &sfEvent) == C.sfFalse {
			// Return nil if the event queue is empty.
			return nil
		}
		event = weEvent(sfEvent)
		if event != nil {
			return event
		}
	}
}

// TODO(u): Figure out a better way to handle the previous cursor position. The
// current way is racy.

// prev represents the previously recorded cursor position.
var prev image.Point

// weEvent returns the we.Event corresponding to the provided SFML event.
func weEvent(sfEvent C.sfEvent) (event we.Event) {
	typ := C.getEventType(sfEvent)
	switch typ {
	// Window events.
	case C.sfEvtClosed:
		return we.Close{}
	case C.sfEvtResized:
		e := C.getSizeEvent(sfEvent)
		event = we.Resize{
			Width:  int(e.width),
			Height: int(e.height),
		}
		return event

	// Keyboard events.
	case C.sfEvtTextEntered:
		e := C.getTextEvent(sfEvent)
		event = we.KeyRune(e.unicode)
		return event
	case C.sfEvtKeyPressed:
		e := C.getKeyEvent(sfEvent)
		event = we.KeyPress{
			Key: weKey(e.code),
			Mod: weMod(e),
		}
		return event
	case C.sfEvtKeyReleased:
		e := C.getKeyEvent(sfEvent)
		event = we.KeyRelease{
			Key: weKey(e.code),
			Mod: weMod(e),
		}
		return event

	// Mouse events.
	case C.sfEvtMouseWheelMoved:
		e := C.getMouseWheelEvent(sfEvent)
		pt := image.Pt(int(e.x), int(e.y))
		event = we.ScrollY{
			Point: pt,
			Off:   int(e.delta),
			Mod:   getMod(),
		}
		return event
	case C.sfEvtMouseButtonPressed:
		e := C.getMouseButtonEvent(sfEvent)
		pt := image.Pt(int(e.x), int(e.y))
		event = we.MousePress{
			Point:  pt,
			Button: weButton(e.button),
			Mod:    getMod(),
		}
		return event
	case C.sfEvtMouseButtonReleased:
		e := C.getMouseButtonEvent(sfEvent)
		pt := image.Pt(int(e.x), int(e.y))
		event = we.MouseRelease{
			Point:  pt,
			Button: weButton(e.button),
			Mod:    getMod(),
		}
		return event
	case C.sfEvtMouseMoved:
		// TODO(u): Implement we.MouseDrag event.
		e := C.getMouseMoveEvent(sfEvent)
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
		log.Printf("window.weEvent: support for SFML event type %d not yet implemented.\n", typ)
	}

	// Ignore event.
	return nil
}
