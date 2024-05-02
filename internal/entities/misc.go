package entities

import "github.com/veandco/go-sdl2/sdl"

type Background struct {
	texture *sdl.Texture
	rect    *sdl.FRect
}

func (b *Background) Draw(r *sdl.Renderer) error {
	return r.CopyF(b.texture, nil, b.rect)
}

func NewBackgroundFactory() BackgroundFactory {
	return BackgroundFactory{}
}

type BackgroundFactory struct {
}

func (b BackgroundFactory) Create(texture *sdl.Texture) *Background {
	return &Background{
		texture: texture,
		rect:    &sdl.FRect{X: 0, Y: 0, W: 800, H: 600},
	}
}
