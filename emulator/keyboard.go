package emulator

type keyboard struct {
	waitKey                  uint8
	onKeyPressedEqualWaitKey func(uint8)
}

func MakeKeyboard() keyboard {
	return keyboard{waitKey: 0, onKeyPressedEqualWaitKey: nil}
}

func (keyboard *keyboard) SetWaitKey(key uint8) {
	keyboard.waitKey = key
}

func (keyboard *keyboard) HandleKeyPressed(key uint8) {
	if key == keyboard.waitKey {
		keyboard.onKeyPressedEqualWaitKey(key)
	}
}
