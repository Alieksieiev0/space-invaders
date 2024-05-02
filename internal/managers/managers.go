package managers

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/veandco/go-sdl2/sdl"
)

type EntityManager interface {
	Spawn(r *sdl.Renderer) error
	Behave(r *sdl.Renderer) error
}

type ProjectileManager interface {
	EntityManager
	AddTarget(entity entities.Entity[float32, *sdl.FRect])
}
