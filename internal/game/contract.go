package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Game interface {
	Start()
}

type Loop interface {
	Run(r *sdl.Renderer) error
}
