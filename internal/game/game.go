package game

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/veandco/go-sdl2/sdl"
)

type Loop interface {
	Run(r *sdl.Renderer) error
}

type KeyboardEventHandler interface {
	HandleKeyboardEvent(event *sdl.KeyboardEvent)
}

type KeyboardEventRenderer interface {
	RenderKeyboardEvent(event *sdl.KeyboardEvent, r *sdl.Renderer)
}

type AfterEventHandler interface {
	HandleAfterEvent()
}

type AfterEventRenderer interface {
	RenderAfterEvent(r *sdl.Renderer)
}

type Shooter interface {
	Fire(r *sdl.Renderer)
	AddTarget(entity entities.Entity[float32, *sdl.FRect])
	InitializeProjectileManager(texture *sdl.Texture)
}

type Player interface {
	entities.MoveableEntity[float32, *sdl.FRect]
	Shooter
	KeyboardEventRenderer
	AfterEventHandler
	AfterEventRenderer
}

type Enemy interface {
	entities.MoveableEntity[float32, *sdl.FRect]
	Shooter
	AfterEventRenderer
}

type EnemyPool interface {
	Shooter
	AfterEventHandler
	AfterEventRenderer
	Enemies() [][]Enemy
	Shooters() []Enemy
	DrawEnemies(r *sdl.Renderer) error
}
