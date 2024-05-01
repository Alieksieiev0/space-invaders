package handlers

import (
	"github.com/veandco/go-sdl2/sdl"
)

func NewProjectileHandler(fire func(), move func()) Handler {
	return &ProjectileHandler{
		fire: fire,
		move: move,
	}
}

type ProjectileHandler struct {
	fire func()
	move func()
}

func (p *ProjectileHandler) HandleKeyboardEvent(event *sdl.KeyboardEvent) {
	switch event.Type {
	case sdl.KEYDOWN:
		switch event.Keysym.Sym {
		case sdl.K_SPACE:
			p.fire()
		}
	}
}

func (p *ProjectileHandler) HandleAfterEvent() {
	p.move()
}
