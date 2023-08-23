package main

import (
	"varian97/go-CHIP8/emulator"
)

func main() {
	cpu := emulator.CPU{}
	cpu.Reset()
	cpu.LoadRom("./roms/Pong.ch8")
}
