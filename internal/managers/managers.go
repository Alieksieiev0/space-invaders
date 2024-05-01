package managers

import "github.com/veandco/go-sdl2/sdl"

type EntityManager interface {
	Spawn(r *sdl.Renderer) error
	Behave(r *sdl.Renderer) error
}
