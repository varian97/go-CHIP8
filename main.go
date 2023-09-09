package main

import (
	"fmt"
	"os"

	"varian97/go-CHIP8/emulator"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: gochip8 <path-to-rom>")
		os.Exit(1)
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	romFile := os.Args[1]

	display := emulator.MakeDisplay()
	keyboard := emulator.MakeKeyboard()
	audio := emulator.MakeAudio(4410, 440)

	cpu := emulator.MakeCPU(display, keyboard, audio)
	err := cpu.LoadRom(romFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = audio.OpenAudio()
	if err != nil {
		fmt.Println(err)
	}
	defer audio.Cleanup()

	display.Init()
	defer display.Cleanup()

	tick0 := sdl.GetTicks64()

	running := true
	for running {
		tick1 := sdl.GetTicks64()
		delta := tick1 - tick0

		if delta > 1000/60 {
			tick0 = tick1

			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					running = false
				case *sdl.KeyboardEvent:
					keycode := t.Keysym.Scancode
					keyboard.HandleKeyPressed(keycode)
				}
			}

			cpu.Cycle()
			display.UpdateSurface()
		}
	}
}
