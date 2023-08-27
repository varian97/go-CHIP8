package emulator

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

type CPU struct {
	instructionsPerCycle int
	memory               [4096]byte
	v                    [16]byte
	i                    uint16
	delayTimer           byte
	soundTimer           byte
	stack                [16]uint16
	stackPointer         int8
	pc                   uint16
	paused               bool
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

	sprites := [80]byte{
		0xf0, 0x90, 0x90, 0x90, 0xf0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xf0, 0x10, 0xf0, 0x80, 0xf0, // 2
		0xf0, 0x10, 0xf0, 0x10, 0xf0, // 3
		0x90, 0x90, 0xf0, 0x10, 0x10, // 4
		0xf0, 0x80, 0xf0, 0x10, 0xf0, // 5
		0xf0, 0x80, 0xf0, 0x90, 0xf0, // 6
		0xf0, 0x10, 0x20, 0x40, 0x40, // 7
		0xf0, 0x90, 0xf0, 0x90, 0xf0, // 8
		0xf0, 0x90, 0xf0, 0x10, 0xf0, // 9
		0xf0, 0x90, 0xf0, 0x90, 0x90, // A
		0xe0, 0x90, 0xe0, 0x90, 0xe0, // B
		0xf0, 0x80, 0x80, 0x80, 0xf0, // C
		0xe0, 0x90, 0x90, 0x90, 0xe0, // D
		0xf0, 0x80, 0xf0, 0x80, 0xf0, // E
		0xf0, 0x80, 0xf0, 0x80, 0x80, // F
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
		cpu.memory[i+0x200] = romContent[i]
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
	firstByte := uint16(cpu.memory[cpu.pc])
	secondByte := uint16(cpu.memory[cpu.pc+1])
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
	case 4:
		cpu.handle4(opcode)
	case 5:
		cpu.handle5(opcode)
	case 6:
		cpu.handle6(opcode)
	case 7:
		cpu.handle7(opcode)
	case 8:
		cpu.handle8(opcode)
	case 9:
		cpu.handle9(opcode)
	case 0xA:
		cpu.handleA(opcode)
	case 0xB:
		cpu.handleB(opcode)
	case 0xC:
		cpu.handleC(opcode)
	case 0xD:
		cpu.handleD(opcode)
	case 0xE:
		cpu.handleE(opcode)
	case 0xF:
		cpu.handleF(opcode)
	default:
		panic(fmt.Errorf("invalid instruction %X", opcode))
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
	if cpu.v[(opcode>>8)&0xf] == byte(opcode&0xff) {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func (cpu *CPU) handle4(opcode uint16) {
	if cpu.v[(opcode>>8)&0xf] != byte(opcode&0xff) {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func (cpu *CPU) handle5(opcode uint16) {
	vx := cpu.v[(opcode>>8)&0xf]
	vy := cpu.v[(opcode>>4)&0xf]
	if vx == vy {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func (cpu *CPU) handle6(opcode uint16) {
	cpu.v[(opcode>>8)&0xf] = byte(opcode & 0xff)
	cpu.pc += 2
}

func (cpu *CPU) handle7(opcode uint16) {
	cpu.v[(opcode>>8)&0xf] += byte(opcode & 0xff)
	cpu.pc += 2
}

func (cpu *CPU) handle8(opcode uint16) {
	x := (opcode >> 8) & 0xf
	y := (opcode >> 4) & 0xf

	switch opcode & 0xf {
	case 0x1:
		cpu.v[x] |= byte(y)
	case 0x2:
		cpu.v[x] &= byte(y)
	case 0x3:
		cpu.v[x] ^= byte(y)
	case 0x4:
		vx := cpu.v[x]
		vy := cpu.v[y]
		if vx > 255-vy {
			cpu.v[0xf] = 1
		} else {
			cpu.v[0xf] = 0
		}
		cpu.v[x] = vx + vy
	case 0x5:
		vx := cpu.v[x]
		vy := cpu.v[y]
		if vx > vy {
			cpu.v[0xf] = 1
		} else {
			cpu.v[0xf] = 0
		}
		cpu.v[x] = vx - vy
	case 0x6:
		vx := cpu.v[x]
		if vx&0x1 > 0 {
			cpu.v[0xf] = 1
		} else {
			cpu.v[0xf] = 0
		}
		cpu.v[x] /= 2
	case 0x7:
		vx := cpu.v[x]
		vy := cpu.v[y]
		if vy > vx {
			cpu.v[0xf] = 1
		} else {
			cpu.v[0xf] = 0
		}
		cpu.v[x] = vy - vx
	case 0xe:
		vx := cpu.v[x]
		if (vx>>7)&0x1 > 0 {
			cpu.v[0xf] = 1
		} else {
			cpu.v[0xf] = 0
		}
		cpu.v[x] *= 2
	}

	cpu.pc += 2
}

func (cpu *CPU) handle9(opcode uint16) {
	if cpu.v[(opcode>>8)&0xf] != cpu.v[(opcode>>4)&0xf] {
		cpu.pc += 2
	}
	cpu.pc += 2
}

func (cpu *CPU) handleA(opcode uint16) {
	cpu.i = (opcode & 0xfff)
	cpu.pc += 2
}

func (cpu *CPU) handleB(opcode uint16) {
	cpu.pc = uint16(cpu.v[(opcode>>8)&0xf] + cpu.v[0])
}

func (cpu *CPU) handleC(opcode uint16) {
	r := byte(rand.Intn(256))
	x := (opcode >> 8) & 0xf
	cpu.v[x] = r & byte((opcode & 0xff))
	cpu.pc += 2
}

func (cpu *CPU) handleD(opcode uint16) {
	// todo: implement drawing to canvas
}

func (cpu *CPU) handleE(opcode uint16) {
	// todo: implement keyboard
}

func (cpu *CPU) handleF(opcode uint16) {
	x := (opcode >> 8) & 0xf
	switch opcode & 0xff {
	case 0x7:
		cpu.v[x] = cpu.delayTimer
	case 0xA:
		// todo: implement wait key pressed
		cpu.paused = true
	case 0x15:
		cpu.delayTimer = cpu.v[x]
	case 0x18:
		cpu.soundTimer = cpu.v[x]
	case 0x1E:
		cpu.i += uint16(cpu.v[x])
	case 0x29:
		cpu.i = uint16(cpu.v[x] * 5)
	case 0x33:
		vx := cpu.v[x]
		cpu.memory[cpu.i] = vx / 100
		cpu.memory[cpu.i+1] = (vx / 10) % 10
		cpu.memory[cpu.i+2] = vx % 10
	case 0x55:
		var i uint16 = 0
		for i = 0; i <= x; i++ {
			cpu.memory[cpu.i+i] = cpu.v[i]
		}
	case 0x65:
		var i uint16 = 0
		for i = 0; i <= x; i++ {
			cpu.v[i] = cpu.memory[cpu.i+i]
		}
	}
	cpu.pc += 2
}
