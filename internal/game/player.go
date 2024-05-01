package game

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/Alieksieiev0/space-invaders/internal/handlers"
	"github.com/veandco/go-sdl2/sdl"
)

func NewGamePlayer(
	entity entities.MoveableEntity[float32, *sdl.FRect],
	handlers []handlers.KeyboardEventHandler,
	afterEventHandlers []handlers.AfterEventHandler,
) Player {
	return &GamePlayer{
		MoveableEntity:     entity,
		handlers:           handlers,
		afterEventHandlers: afterEventHandlers,
	}
}

type GamePlayer struct {
	entities.MoveableEntity[float32, *sdl.FRect]
	targets            []entities.Entity[float32, *sdl.FRect]
	handlers           []handlers.KeyboardEventHandler
	afterEventHandlers []handlers.AfterEventHandler
}

func (g *GamePlayer) HandleKeyboardEvent(event *sdl.KeyboardEvent) {
	for _, h := range g.handlers {
		h.HandleKeyboardEvent(event)
	}
}

func (g *GamePlayer) HandleAfterEvent() {
	for _, h := range g.afterEventHandlers {
		h.HandleAfterEvent()
	}
}

func (g *GamePlayer) AddTarget(entity entities.Entity[float32, *sdl.FRect]) {
	g.targets = append(g.targets, entity)
}

func (g *GamePlayer) Targets() []entities.Entity[float32, *sdl.FRect] {
	return g.targets
}
