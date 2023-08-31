package emulator

import (
	"github.com/veandco/go-sdl2/sdl"
)

var USER_KEY_TO_SDL_KEY = map[uint8]sdl.Scancode{
	0:   sdl.SCANCODE_X,
	1:   sdl.SCANCODE_1,
	2:   sdl.SCANCODE_2,
	3:   sdl.SCANCODE_3,
	4:   sdl.SCANCODE_Q,
	5:   sdl.SCANCODE_W,
	6:   sdl.SCANCODE_E,
	7:   sdl.SCANCODE_A,
	8:   sdl.SCANCODE_S,
	9:   sdl.SCANCODE_D,
	0xa: sdl.SCANCODE_Z,
	0xb: sdl.SCANCODE_C,
	0xc: sdl.SCANCODE_4,
	0xd: sdl.SCANCODE_R,
	0xe: sdl.SCANCODE_F,
	0xf: sdl.SCANCODE_V,
}

var SDL_KEY_TO_USER_KEY = map[sdl.Scancode]uint8{
	sdl.SCANCODE_X: 0,
	sdl.SCANCODE_1: 1,
	sdl.SCANCODE_2: 2,
	sdl.SCANCODE_3: 3,
	sdl.SCANCODE_Q: 4,
	sdl.SCANCODE_W: 5,
	sdl.SCANCODE_E: 6,
	sdl.SCANCODE_A: 7,
	sdl.SCANCODE_S: 8,
	sdl.SCANCODE_D: 9,
	sdl.SCANCODE_Z: 0xa,
	sdl.SCANCODE_C: 0xb,
	sdl.SCANCODE_4: 0xc,
	sdl.SCANCODE_R: 0xd,
	sdl.SCANCODE_F: 0xe,
	sdl.SCANCODE_V: 0xf,
}

type keyboard struct {
	waitKey                  bool
	onKeyPressedEqualWaitKey func(uint8)
}

func MakeKeyboard() *keyboard {
	return &keyboard{waitKey: false, onKeyPressedEqualWaitKey: nil}
}

func (keyboard *keyboard) SetWaitKey(b bool) {
	keyboard.waitKey = b
}

func (keyboard *keyboard) SetOnKeyPressedEqualWaitKey(callback func(uint8)) {
	keyboard.onKeyPressedEqualWaitKey = callback
}

func (keyboard keyboard) IsKeyPressed(userKey uint8) bool {
	keyboardState := sdl.GetKeyboardState()
	pressedKey, ok := USER_KEY_TO_SDL_KEY[userKey]
	if ok {
		return keyboardState[pressedKey] == 1
	}
	return false
}

func (keyboard *keyboard) HandleKeyPressed(key sdl.Scancode) {
	userKey, ok := SDL_KEY_TO_USER_KEY[key]
	if ok && keyboard.waitKey {
		keyboard.SetWaitKey(false)
		keyboard.onKeyPressedEqualWaitKey(userKey)
	}
}
