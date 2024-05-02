package entities

import "github.com/veandco/go-sdl2/sdl"

type EntityType int64

const (
	Player EntityType = iota
	Enemy
	Obstacle
	Missile
)

type StaticEntity interface {
	Draw(renderer *sdl.Renderer) error
}

type Entity[T int32 | float32, V *sdl.Rect | *sdl.FRect] interface {
	StaticEntity
	Dimensions() (x, y, w, h T)
	Rotate(renderer *sdl.Renderer, angle float64) error

	// MAKES COPY!
	Rect() V

	LoadTexture(texture *sdl.Texture)
	Intersect(entity Entity[T, V]) bool
	Destroy()
	IsAlive() bool
}

type Movement[T int32 | float32] interface {
	StepX() T
	StepY() T
	Move(x T, y T)
}

type MoveableEntity[T int32 | float32, V *sdl.Rect | *sdl.FRect] interface {
	Entity[T, V]
	Movement[T]
}
