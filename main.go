package main

import (
	"varian97/go-CHIP8/emulator"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	display := emulator.MakeDisplay()
	cpu := emulator.MakeCPU(display)

	cpu.LoadRom("./roms/Pong.ch8")

	display.Init()
	defer display.Cleanup()

	running := true
	for running {
		cpu.Cycle()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
	}
}
