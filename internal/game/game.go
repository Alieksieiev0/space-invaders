package game

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/Alieksieiev0/space-invaders/internal/handlers"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func NewGameManager() *GameManager {
	return &GameManager{}
}

type GameManager struct {
}

func (g *GameManager) Start() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"test",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800,
		600,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return err
	}
	defer window.Destroy()
	r, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	bg := entities.NewBackgroundFactory().CreateWhite()
	if err = bg.Draw(r); err != nil {
		return err
	}

	if err = r.Clear(); err != nil {
		return err
	}
	ship_img, err := img.Load("assets/ship.png")
	if err != nil {
		return err
	}
	ship, err := r.CreateTextureFromSurface(ship_img)
	if err != nil {
		return err
	}

	missile_img, err := img.Load("assets/sprite_1.png")
	if err != nil {
		return err
	}
	missile, err := r.CreateTextureFromSurface(missile_img)
	if err != nil {
		return err
	}
	frect := entities.NewFloatRectFactory().
		Create(ship, entities.Player, 375, 275, 50, 50)
	projectile := entities.NewBidirectionalProjectile(frect, missile, 10, 20)

	movementHandler := handlers.NewHorizontalMovementHandler(
		func() {
			frect.Move(-frect.StepX(), 0)
		},
		func() {
			frect.Move(frect.StepX(), 0)
		},
	)
	projectileHandler := handlers.NewProjectileHandler(
		func() {
			projectile.Fire(r)
		},
		func() {
			projectile.Move(r)
		},
	)

	player := entities.NewPlayerFactory().
		Create(frect, []entities.KeyboardEventHandler{projectileHandler}, []entities.AfterEventHandler{movementHandler, projectileHandler})
	if err := player.Draw(r); err != nil {
		return err
	}
	r.Present()

	if err := bg.Draw(r); err != nil {
		return err
	}

	loop := NewGameLoop(player, bg)
	return loop.Run(r)
}
