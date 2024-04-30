package entities

import (
	"github.com/Alieksieiev0/space-invaders/internal/mechanics"
	"github.com/veandco/go-sdl2/sdl"
)

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

type Entity[T int32 | float32] interface {
	StaticEntity
	Dimensions() (x, y, w, h T)
}

type KeyboardEventHandler interface {
	HandleKeyboardEvent(event *sdl.KeyboardEvent)
}

type AfterEventHandler interface {
	HandleAfterEvent()
}

type Handler interface {
	KeyboardEventHandler
	AfterEventHandler
}

type MoveableEntity[T int32 | float32] interface {
	Entity[T]
	mechanics.Movement[T]
}

type PlayerEntity interface {
	MoveableEntity[float32]
	KeyboardEventHandler
	AfterEventHandler
}

type Projectile interface {
	Fire(r *sdl.Renderer)
	Move(r *sdl.Renderer)
}
