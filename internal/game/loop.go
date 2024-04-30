package game

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	entity "github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	fps        = 60
	frameDelay = 1000 / fps
)

func NewGameLoop(player entities.PlayerEntity, bg entities.StaticEntity) Loop {
	return &GameLoop{player: player, bg: bg}
}

type GameLoop struct {
	player entity.PlayerEntity
	bg     entity.StaticEntity
}

func (g *GameLoop) Run(r *sdl.Renderer) error {
	var frameStart uint64
	var frameTime uint64
	for {
		frameStart = sdl.GetTicks64()

		if err := r.Clear(); err != nil {
			return err
		}
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch e := event.(type) {
			case *sdl.KeyboardEvent:
				g.player.HandleKeyboardEvent(e)
			case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				return nil
			}
		}

		g.player.HandleAfterEvent()
		if err := g.player.Draw(r); err != nil {
			return err
		}
		if err := g.bg.Draw(r); err != nil {
			return err
		}
		r.Present()

		frameTime = sdl.GetTicks64() - frameStart
		if frameDelay > frameTime {
			sdl.Delay(uint32(frameDelay - frameTime))
		}

		// if err := r.Clear(); err != nil {
		// 	return err
		// }
		// g.player.Update()
		// if err := g.player.Draw(r); err != nil {
		// 	return err
		// }
		// if err := g.bg.Draw(r); err != nil {
		// 	return err
		// }
		// r.Present()

	}
}
