package window

// #include <SFML/Window.h>
import "C"

import (
	"log"

	"github.com/mewspring/we"
)

// getMod returns the active keyboard modifiers.
func getMod() we.Mod {
	var mod we.Mod
	if C.sfKeyboard_isKeyPressed(C.sfKeyLShift) == C.sfTrue ||
		C.sfKeyboard_isKeyPressed(C.sfKeyRShift) == C.sfTrue {
		mod |= we.ModShift
	}
	if C.sfKeyboard_isKeyPressed(C.sfKeyLControl) == C.sfTrue ||
		C.sfKeyboard_isKeyPressed(C.sfKeyRControl) == C.sfTrue {
		mod |= we.ModControl
	}
	if C.sfKeyboard_isKeyPressed(C.sfKeyLAlt) == C.sfTrue ||
		C.sfKeyboard_isKeyPressed(C.sfKeyRAlt) == C.sfTrue {
		mod |= we.ModAlt
	}
	if C.sfKeyboard_isKeyPressed(C.sfKeyLSystem) == C.sfTrue ||
		C.sfKeyboard_isKeyPressed(C.sfKeyRSystem) == C.sfTrue {
		mod |= we.ModSuper
	}
	return mod
}

// weMod returns the we.Mod corresponding to the provided SFML modifiers.
func weMod(e C.sfKeyEvent) we.Mod {
	var mod we.Mod
	if e.shift == C.sfTrue {
		mod |= we.ModShift
	}
	if e.control == C.sfTrue {
		mod |= we.ModControl
	}
	if e.alt == C.sfTrue {
		mod |= we.ModAlt
	}
	if e.system == C.sfTrue {
		mod |= we.ModSuper
	}
	return mod
}

// weKey returns the we.Key corresponding to the provided SFML key code.
func weKey(code C.sfKeyCode) we.Key {
	switch code {
	case C.sfKeyA:
		return we.KeyA
	case C.sfKeyB:
		return we.KeyB
	case C.sfKeyC:
		return we.KeyC
	case C.sfKeyD:
		return we.KeyD
	case C.sfKeyE:
		return we.KeyE
	case C.sfKeyF:
		return we.KeyF
	case C.sfKeyG:
		return we.KeyG
	case C.sfKeyH:
		return we.KeyH
	case C.sfKeyI:
		return we.KeyI
	case C.sfKeyJ:
		return we.KeyJ
	case C.sfKeyK:
		return we.KeyK
	case C.sfKeyL:
		return we.KeyL
	case C.sfKeyM:
		return we.KeyM
	case C.sfKeyN:
		return we.KeyN
	case C.sfKeyO:
		return we.KeyO
	case C.sfKeyP:
		return we.KeyP
	case C.sfKeyQ:
		return we.KeyQ
	case C.sfKeyR:
		return we.KeyR
	case C.sfKeyS:
		return we.KeyS
	case C.sfKeyT:
		return we.KeyT
	case C.sfKeyU:
		return we.KeyU
	case C.sfKeyV:
		return we.KeyV
	case C.sfKeyW:
		return we.KeyW
	case C.sfKeyX:
		return we.KeyX
	case C.sfKeyY:
		return we.KeyY
	case C.sfKeyZ:
		return we.KeyZ
	case C.sfKeyNum0:
		return we.Key0
	case C.sfKeyNum1:
		return we.Key1
	case C.sfKeyNum2:
		return we.Key2
	case C.sfKeyNum3:
		return we.Key3
	case C.sfKeyNum4:
		return we.Key4
	case C.sfKeyNum5:
		return we.Key5
	case C.sfKeyNum6:
		return we.Key6
	case C.sfKeyNum7:
		return we.Key7
	case C.sfKeyNum8:
		return we.Key8
	case C.sfKeyNum9:
		return we.Key9
	case C.sfKeyEscape:
		return we.KeyEscape
	case C.sfKeyLControl:
		return we.KeyLeftControl
	case C.sfKeyLShift:
		return we.KeyLeftShift
	case C.sfKeyLAlt:
		return we.KeyLeftAlt
	case C.sfKeyLSystem:
		return we.KeyLeftSuper
	case C.sfKeyRControl:
		return we.KeyRightControl
	case C.sfKeyRShift:
		return we.KeyRightShift
	case C.sfKeyRAlt:
		return we.KeyRightAlt
	case C.sfKeyRSystem:
		return we.KeyRightSuper
	case C.sfKeyMenu:
		return we.KeyMenu
	case C.sfKeyLBracket:
		return we.KeyLeftBracket
	case C.sfKeyRBracket:
		return we.KeyRightBracket
	case C.sfKeySemiColon:
		return we.KeySemicolon
	case C.sfKeyComma:
		return we.KeyComma
	case C.sfKeyPeriod:
		return we.KeyPeriod
	//case C.sfKeyQuote:
	//	return we.KeyXXX
	case C.sfKeySlash:
		return we.KeySlash
	case C.sfKeyBackSlash:
		return we.KeyBackslash
	//case C.sfKeyTilde:
	//	return we.KeyXXX
	case C.sfKeyEqual:
		return we.KeyEqual
	case C.sfKeyDash:
		return we.KeyMinus
	case C.sfKeySpace:
		return we.KeySpace
	case C.sfKeyReturn:
		return we.KeyEnter
	case C.sfKeyBack:
		return we.KeyBackspace
	case C.sfKeyTab:
		return we.KeyTab
	case C.sfKeyPageUp:
		return we.KeyPageUp
	case C.sfKeyPageDown:
		return we.KeyPageDown
	case C.sfKeyEnd:
		return we.KeyEnd
	case C.sfKeyHome:
		return we.KeyHome
	case C.sfKeyInsert:
		return we.KeyInsert
	case C.sfKeyDelete:
		return we.KeyDelete
	case C.sfKeyAdd:
		return we.KeyKpAdd
	case C.sfKeySubtract:
		return we.KeyKpSubtract
	case C.sfKeyMultiply:
		return we.KeyKpMultiply
	case C.sfKeyDivide:
		return we.KeyKpDivide
	case C.sfKeyLeft:
		return we.KeyLeft
	case C.sfKeyRight:
		return we.KeyRight
	case C.sfKeyUp:
		return we.KeyUp
	case C.sfKeyDown:
		return we.KeyDown
	case C.sfKeyNumpad0:
		return we.KeyKp0
	case C.sfKeyNumpad1:
		return we.KeyKp1
	case C.sfKeyNumpad2:
		return we.KeyKp2
	case C.sfKeyNumpad3:
		return we.KeyKp3
	case C.sfKeyNumpad4:
		return we.KeyKp4
	case C.sfKeyNumpad5:
		return we.KeyKp5
	case C.sfKeyNumpad6:
		return we.KeyKp6
	case C.sfKeyNumpad7:
		return we.KeyKp7
	case C.sfKeyNumpad8:
		return we.KeyKp8
	case C.sfKeyNumpad9:
		return we.KeyKp9
	case C.sfKeyF1:
		return we.KeyF1
	case C.sfKeyF2:
		return we.KeyF2
	case C.sfKeyF3:
		return we.KeyF3
	case C.sfKeyF4:
		return we.KeyF4
	case C.sfKeyF5:
		return we.KeyF5
	case C.sfKeyF6:
		return we.KeyF6
	case C.sfKeyF7:
		return we.KeyF7
	case C.sfKeyF8:
		return we.KeyF8
	case C.sfKeyF9:
		return we.KeyF9
	case C.sfKeyF10:
		return we.KeyF10
	case C.sfKeyF11:
		return we.KeyF11
	case C.sfKeyF12:
		return we.KeyF12
	case C.sfKeyF13:
		return we.KeyF13
	case C.sfKeyF14:
		return we.KeyF14
	case C.sfKeyF15:
		return we.KeyF15
	case C.sfKeyPause:
		return we.KeyPause
	}

	// Unknown key code.
	log.Printf("window.weKey: unknown key code %d", code)
	return 0
}
