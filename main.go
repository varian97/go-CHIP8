package main

import (
	"varian97/go-CHIP8/emulator"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	display := emulator.MakeDisplay()
	keyboard := emulator.MakeKeyboard()
	// audio := emulator.MakeAudio(440, 4400)
	cpu := emulator.MakeCPU(display, keyboard)

	// TODO: REMOVE
	// cpu.LoadRom("./roms/1-chip8-logo.ch8")
	// cpu.LoadRom("./roms/2-ibm-logo.ch8")
	// cpu.LoadRom("./roms/3-corax+.ch8")
	// cpu.LoadRom("./roms/4-flags.ch8")
	// cpu.LoadRom("./roms/5-quirks.ch8")
	// cpu.LoadRom("./roms/6-keypad.ch8")
	cpu.LoadRom("./roms/SpaceInvaders.ch8")

	display.Init()

	// err := audio.OpenAudio()
	// if err != nil {
	// 	log.Println("Failed to open audio ", err)
	// }

	defer sdl.Quit()
	defer display.Cleanup()
	// defer audio.Cleanup()

	tick0 := sdl.GetTicks64()

	running := true
	for running {
		tick1 := sdl.GetTicks64()
		delta := tick1 - tick0

		if delta > 1000/60 {
			tick0 = tick1

			cpu.Cycle()

			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					running = false
				case *sdl.KeyboardEvent:
					keycode := t.Keysym.Scancode
					keyboard.HandleKeyPressed(keycode)
				}
			}
		}
	}
}
