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

	romFile := os.Args[1]

	display := emulator.MakeDisplay()
	keyboard := emulator.MakeKeyboard()

	cpu := emulator.MakeCPU(display, keyboard)
	err := cpu.LoadRom(romFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	display.Init()

	defer sdl.Quit()
	defer display.Cleanup()

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
