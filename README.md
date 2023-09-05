# go-CHIP8
CHIP-8 Interpreter using Go Programming Language

## Table of contents

  * [What is CHIP-8](#what-is-chip-8)
  * [Installation](#installation)
    * [Binary Download](#binary-download)
    * [Build from source](#build-from-source)
  * [Screenshots](#screenshots)
  * [Testing and Correctness](#testing-and-correctness)
  * [Todo](#todo)
  * [References](#references)

## What is CHIP-8
CHIP-8 is an interpreted minimalist programming language that was designed by Joseph Weisbecker in the 1970s for use on the RCA COSMAC VIP computer.

So CHIP-8 is actually a language (often mislabelled as emulator) and this project is about interpreting that language (CHIP-8 Interpreter).

However I also made a mistake at the start of this project, I also thought that CHIP-8 is an emulator. The package in this project even named `emulator`.
Sadly, I am too lazy to change it.

## Installation

### Binary Download
1. Download the binary from [Release Page](https://github.com/varian97/go-CHIP8/releases)
2. Run the program `./gochip8 <path to rom file>`

### Build from source
1. Download and Install [Go](https://go.dev/doc/install)
2. Follow instruction on [go-sdl2](https://github.com/veandco/go-sdl2) to setup SDL on your machine
3. `cd` to the project directory and `go mod tidy`
4. Build using `go build -o gochip8`
5. The binary `gochip8` should be created. Run the program `./gochip8 <path to rom file>`

## Screenshots
![Screenshot from 2023-09-03 22-00-58](https://github.com/varian97/go-CHIP8/assets/15712645/91c06c50-52a1-4e0d-96a9-00b8fd58a34f)

![Screenshot from 2023-09-03 22-01-32](https://github.com/varian97/go-CHIP8/assets/15712645/6f6142a9-7ae7-4cf0-ae25-2a6aa7c6b65d)


## Testing and Correctness
This project use [chip8-test-suite](https://github.com/Timendus/chip8-test-suite) for testing.
Here is the results:
| Test                  | Result        |
| -------------         | ------------- |
| CHIP-8 splash screen  | ok(1/1)       |
| IBM logo              | ok(1/1)       |
| Corax+ opcode test    | ok(22/22)     |
| Flags test            | ok(14/14)     |
| Quirks test           | not ok(5/6)   |

## Todo
- [ ] 100% Test passed
- [ ] Add menu bar in the UI
- [ ] Add support for debugging

## References
- http://devernay.free.fr/hacks/chip8/C8TECH10.HTM
- https://github.com/Timendus/chip8-test-suite
- https://tobiasvl.github.io/blog/write-a-chip-8-emulator/
- https://github.com/mattmikolay/chip-8/wiki/Mastering-CHIP%E2%80%908
