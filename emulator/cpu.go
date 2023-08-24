package emulator

import (
	"fmt"
	"io"
	"os"
)

type CPU struct {
	instructionsPerCycle int
	memory [4096]byte
	v [16]byte
	i uint16
	delayTimer byte
	soundTimer byte
	stack [16]uint16
	stackPointer int8
	pc uint16
	paused bool
}

func (cpu *CPU) Reset() {
	cpu.pc = 0x200
	cpu.delayTimer = 0
	cpu.soundTimer = 0
	cpu.i = 0
	cpu.instructionsPerCycle = 10
	cpu.paused = false
	cpu.stackPointer = -1

	for i := 0; i < 16; i++ {
		cpu.stack[i] = 0
		cpu.v[i] = 0
	}

	sprites := [80]byte {
		0xf0, 0x90, 0x90, 0x90, 0xf0,  // 0
		0x20, 0x60, 0x20, 0x20, 0x70,  // 1
		0xf0, 0x10, 0xf0, 0x80, 0xf0,  // 2
		0xf0, 0x10, 0xf0, 0x10, 0xf0,  // 3
		0x90, 0x90, 0xf0, 0x10, 0x10,  // 4
		0xf0, 0x80, 0xf0, 0x10, 0xf0,  // 5
		0xf0, 0x80, 0xf0, 0x90, 0xf0,  // 6
		0xf0, 0x10, 0x20, 0x40, 0x40,  // 7
		0xf0, 0x90, 0xf0, 0x90, 0xf0,  // 8
		0xf0, 0x90, 0xf0, 0x10, 0xf0,  // 9
		0xf0, 0x90, 0xf0, 0x90, 0x90,  // A
		0xe0, 0x90, 0xe0, 0x90, 0xe0,  // B
		0xf0, 0x80, 0x80, 0x80, 0xf0,  // C
		0xe0, 0x90, 0x90, 0x90, 0xe0,  // D
		0xf0, 0x80, 0xf0, 0x80, 0xf0,  // E
		0xf0, 0x80, 0xf0, 0x80, 0x80,  // F
	}
	for i := 0; i < cap(cpu.memory); i++ {
		if i < len(sprites) {
			cpu.memory[i] = sprites[i]
		} else {
			cpu.memory[i] = 0
		}
	}
}

func (cpu *CPU) LoadRom(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	romContent, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	for i := 0; i < len(romContent); i++ {
		cpu.memory[i + 0x200] = romContent[i]
	}

	return nil
}

func (cpu *CPU) Cycle() {
	for i := 0; i < cpu.instructionsPerCycle; i++ {
		if !cpu.paused {
			opcode := cpu.fetchInstruction()
			cpu.runInstruction(opcode)
		}
	}

	// todo: update timer here

	// todo: play sound here
}

func (cpu CPU) fetchInstruction() uint16 {
	firstByte  := uint16(cpu.memory[cpu.pc])
	secondByte := uint16(cpu.memory[cpu.pc + 1])
	return (firstByte << 8) | secondByte
}

func (cpu *CPU) runInstruction(opcode uint16) {
	prefix := opcode >> 12
	switch prefix {
	case 0:
		cpu.handle0(opcode)
	case 1:
		cpu.handle1(opcode)
	case 2:
		cpu.handle2(opcode)
	case 3:
		cpu.handle3(opcode)
	default:
		panic(fmt.Errorf("instruction %X not yet implemented", opcode))
	}
}

func (cpu *CPU) handle0(opcode uint16) {
	switch opcode {
	case 0x00e0:
		// todo: clear display here
		cpu.pc += 2
	case 0x00ee:
		cpu.pc = cpu.stack[cpu.stackPointer]
		cpu.stackPointer -= 1
	}
}

func (cpu *CPU) handle1(opcode uint16) {
	cpu.pc = opcode & 0xfff
}

func (cpu *CPU) handle2(opcode uint16) {
	cpu.stackPointer += 1
	cpu.stack[cpu.stackPointer] = cpu.pc + 2
	cpu.pc = opcode & 0xfff
}

func (cpu *CPU) handle3(opcode uint16) {
	if cpu.v[(opcode >> 8) & 0xf] == byte(opcode & 0xff) {
		cpu.pc += 2
	}
	cpu.pc += 2
}
