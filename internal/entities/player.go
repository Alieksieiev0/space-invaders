package entities

import "github.com/veandco/go-sdl2/sdl"

type BasePlayer struct {
	MoveableEntity[float32]
	handlers           []KeyboardEventHandler
	afterEventHandlers []AfterEventHandler
}

func (b *BasePlayer) HandleKeyboardEvent(event *sdl.KeyboardEvent) {
	for _, h := range b.handlers {
		h.HandleKeyboardEvent(event)
	}
}

func (b *BasePlayer) HandleAfterEvent() {
	for _, h := range b.afterEventHandlers {
		h.HandleAfterEvent()
	}
}

func NewPlayerFactory() PlayerFactory {
	return PlayerFactory{}
}

type PlayerFactory struct {
}

func (p PlayerFactory) Create(
	entity MoveableEntity[float32],
	handlers []KeyboardEventHandler,
	afterEventHandlers []AfterEventHandler,
) PlayerEntity {
	return &BasePlayer{
		MoveableEntity:     entity,
		handlers:           handlers,
		afterEventHandlers: afterEventHandlers,
	}
}
