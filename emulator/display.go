package emulator

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	ROW    int32 = 32
	COL    int32 = 64
	SCALE  int32 = 10
	HEIGHT int32 = ROW * SCALE
	WIDTH  int32 = COL * SCALE
)

type display struct {
	window         *sdl.Window
	surface        *sdl.Surface
	logicalSurface [HEIGHT][WIDTH]bool
}

func MakeDisplay() *display {
	display := &display{}
	return display
}

func (display *display) Init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("CHIP-8", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		int32(WIDTH), int32(HEIGHT), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	display.window = window
	display.surface = surface
	display.Clear()
}

func (display display) Cleanup() {
	display.surface.Free()
	display.window.Destroy()
}

func (display *display) Clear() {
	for i := int32(0); i < HEIGHT; i++ {
		for j := int32(0); j < WIDTH; j++ {
			display.logicalSurface[i][j] = false
		}
	}

	color := sdl.Color{R: 36, G: 38, B: 39, A: 1}
	pixelColor := sdl.MapRGBA(display.surface.Format, color.R, color.G, color.B, color.A)
	display.surface.FillRect(nil, pixelColor)

	display.window.UpdateSurface()
}

func (display *display) SetPixel(x, y int32) bool {
	rect := sdl.Rect{X: x * SCALE, Y: y * SCALE, W: SCALE, H: SCALE}
	color := sdl.Color{R: 36, G: 38, B: 39, A: 1}
	isUnset := false

	if display.logicalSurface[x][y] {
		display.logicalSurface[x][y] = false
		isUnset = true
	} else {
		display.logicalSurface[x][y] = true
		color = sdl.Color{R: 255, G: 255, B: 255, A: 1}
	}

	pixelColor := sdl.MapRGBA(display.surface.Format, color.R, color.G, color.B, color.A)
	display.surface.FillRect(&rect, pixelColor)

	display.window.UpdateSurface()
	return isUnset
}
