package handlers

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/veandco/go-sdl2/sdl"
)

func NewHorizontalMovementHandler(left func(), right func()) entities.AfterEventHandler {
	return &HorizontalMovementHandler{
		moveLeft:  left,
		moveRight: right,
	}
}

type HorizontalMovementHandler struct {
	moveLeft  func()
	moveRight func()
}

func (h *HorizontalMovementHandler) HandleAfterEvent() {
	state := sdl.GetKeyboardState()
	if state[sdl.SCANCODE_LEFT] == 1 {
		h.moveLeft()
	} else if state[sdl.SCANCODE_RIGHT] == 1 {
		h.moveRight()
	}
}
