package entities

import "github.com/veandco/go-sdl2/sdl"

func NewColorFactory() ColorFactory {
	return ColorFactory{}
}

type ColorFactory struct {
}

func (c ColorFactory) CreateRed() sdl.Color {
	return sdl.Color{R: 255, G: 0, B: 0, A: 1}
}

func (c ColorFactory) CreateBlue() sdl.Color {
	return sdl.Color{R: 0, G: 0, B: 255, A: 1}
}

type Background struct {
	color sdl.Color
}

func (b *Background) Draw(r *sdl.Renderer) error {
	return r.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
}

func NewBackgroundFactory() BackgroundFactory {
	return BackgroundFactory{}
}

type BackgroundFactory struct {
}

func (b BackgroundFactory) CreateWhite() StaticEntity {
	return &Background{
		color: sdl.Color{R: 255, G: 255, B: 255, A: 0},
	}
}
