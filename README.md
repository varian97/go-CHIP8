# go-CHIP8

## What is CHIP-8
CHIP-8 is an interpreted minimalist programming language that was designed by Joseph Weisbecker in the 1970s for use on the RCA COSMAC VIP computer.
So CHIP-8 is actually an language (often mislabelled as emulator) and this project is about interpreting that language (CHIP-8 Interpreter).

However I also made a mistake at the start of this project, I too, thought that CHIP-8 is an emulator. The package in this project even named `emulator`.
Sadly, I am too lazy to change it.

## Build from source
1. Download and Install [Go](https://go.dev/doc/install)
2. Follow instruction on [go-sdl2](https://github.com/veandco/go-sdl2) to setup SDL on your machine
3. `cd` to the project directory and `go mod tidy`
4. Build using `go build`
5. The binary `go-CHIP8` should be created. Run the program `./go-CHIP8 <path to rom file>`

## References
- http://devernay.free.fr/hacks/chip8/C8TECH10.HTM
- https://github.com/Timendus/chip8-test-suite
- https://tobiasvl.github.io/blog/write-a-chip-8-emulator/
- https://github.com/mattmikolay/chip-8/wiki/Mastering-CHIP%E2%80%908
