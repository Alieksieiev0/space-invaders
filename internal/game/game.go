package game

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/Alieksieiev0/space-invaders/internal/handlers"
	"github.com/veandco/go-sdl2/sdl"
)

type Loop interface {
	Run(r *sdl.Renderer) error
}

type Shooter interface {
	AddTarget(entity entities.Entity[float32, *sdl.FRect])
	Targets() []entities.Entity[float32, *sdl.FRect]
}

type Player interface {
	Shooter
	entities.MoveableEntity[float32, *sdl.FRect]
	handlers.KeyboardEventHandler
	handlers.AfterEventHandler
}

type Enemy interface {
	Shooter
	entities.MoveableEntity[float32, *sdl.FRect]
	handlers.AfterEventHandler
}

type EnemyPool interface {
	Shooter
	handlers.AfterEventHandler
	Enemies() []Enemy
	DrawEnemies(r *sdl.Renderer) error
}
