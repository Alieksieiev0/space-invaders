package game

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/Alieksieiev0/space-invaders/internal/managers"
	"github.com/veandco/go-sdl2/sdl"
)

func NewGamePlayer(entity entities.MoveableEntity[float32, *sdl.FRect]) Player {
	return &GamePlayer{
		MoveableEntity: entity,
	}
}

type GamePlayer struct {
	entities.MoveableEntity[float32, *sdl.FRect]
	projectileManager managers.ProjectileManager
}

func (g *GamePlayer) RenderKeyboardEvent(event *sdl.KeyboardEvent, r *sdl.Renderer) {
	switch event.Type {
	case sdl.KEYDOWN:
		switch event.Keysym.Sym {
		case sdl.K_SPACE:
			g.projectileManager.Spawn(r)
		}
	}
}

func (g *GamePlayer) HandleAfterEvent() {
	state := sdl.GetKeyboardState()
	if state[sdl.SCANCODE_LEFT] == 1 {
		g.Move(-g.StepX(), 0)
	} else if state[sdl.SCANCODE_RIGHT] == 1 {
		g.Move(g.StepX(), 0)
	}
}

func (g *GamePlayer) RenderAfterEvent(r *sdl.Renderer) {
	g.projectileManager.Behave(r)
}

func (g *GamePlayer) AddTarget(entity entities.Entity[float32, *sdl.FRect]) {
	g.projectileManager.AddTarget(entity)
}

func (g *GamePlayer) Fire(r *sdl.Renderer) {
	if g.projectileManager != nil {
		g.projectileManager.Spawn(r)
	}
}

func (g *GamePlayer) InitializeProjectileManager(texture *sdl.Texture) {
	g.projectileManager = managers.NewBidirectionalProjectileFactory().
		CreateInNegativeDirection(g, texture)
}
