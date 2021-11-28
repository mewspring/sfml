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
// sfMouseWheelScrollEvent getMouseWheelScrollEvent(sfEvent e) {
//    return e.mouseWheelScroll;
// }
import "C"

import (
	"image"
	"log"

	"github.com/mewspring/we"
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
func (win Window) PollEvent() we.Event {
	// Poll the event queue until we locate a non-nil event or the queue is
	// empty.
	var sfEvent C.sfEvent
	for {
		if C.sfRenderWindow_pollEvent(win.win, &sfEvent) == C.sfFalse {
			// Return nil if the event queue is empty.
			return nil
		}
		if event := weEvent(sfEvent); event != nil {
			return event
		}
	}
}

// TODO(u): Figure out a better way to handle the previous cursor position. The
// current way is racy.

// prev represents the previously recorded cursor position.
var prev image.Point

// weEvent returns the corresponding we.Event for the provided SFML event or nil
// if no such event exists.
func weEvent(sfEvent C.sfEvent) we.Event {
	typ := C.getEventType(sfEvent)
	switch typ {
	// Window events.
	case C.sfEvtClosed:
		return we.Close{}
	case C.sfEvtResized:
		e := C.getSizeEvent(sfEvent)
		return we.Resize{
			Width:  int(e.width),
			Height: int(e.height),
		}

	// Keyboard events.
	case C.sfEvtTextEntered:
		e := C.getTextEvent(sfEvent)
		return we.KeyRune(e.unicode)
	case C.sfEvtKeyPressed:
		e := C.getKeyEvent(sfEvent)
		return we.KeyPress{
			Key: weKey(e.code),
			Mod: weMod(e),
		}
	case C.sfEvtKeyReleased:
		e := C.getKeyEvent(sfEvent)
		return we.KeyRelease{
			Key: weKey(e.code),
			Mod: weMod(e),
		}

	// Mouse events.
	case C.sfEvtMouseWheelMoved:
		e := C.getMouseWheelScrollEvent(sfEvent)
		pt := image.Pt(int(e.x), int(e.y))
		switch e.wheel {
		case C.sfMouseHorizontalWheel:
			return we.ScrollX{
				Point: pt,
				Off:   int(e.delta),
				Mod:   getMod(),
			}
		case C.sfMouseVerticalWheel:
			return we.ScrollY{
				Point: pt,
				Off:   int(e.delta),
				Mod:   getMod(),
			}
		default:
			log.Printf("support for mouse wheel %v not yet implemented", e.wheel)
			return nil
		}
	case C.sfEvtMouseButtonPressed:
		e := C.getMouseButtonEvent(sfEvent)
		pt := image.Pt(int(e.x), int(e.y))
		return we.MousePress{
			Point:  pt,
			Button: weButton(e.button),
			Mod:    getMod(),
		}
	case C.sfEvtMouseButtonReleased:
		e := C.getMouseButtonEvent(sfEvent)
		pt := image.Pt(int(e.x), int(e.y))
		return we.MouseRelease{
			Point:  pt,
			Button: weButton(e.button),
			Mod:    getMod(),
		}
	case C.sfEvtMouseMoved:
		// TODO(u): Implement we.MouseDrag event.
		e := C.getMouseMoveEvent(sfEvent)
		pt := image.Pt(int(e.x), int(e.y))
		event := we.MouseMove{
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
		log.Printf("window.weEvent: support for SFML event type %d not yet implemented", typ)
		return nil
	}
}
