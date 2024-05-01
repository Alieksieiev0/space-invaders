package game

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/Alieksieiev0/space-invaders/internal/handlers"
	"github.com/Alieksieiev0/space-invaders/internal/managers"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowWidth  = 800
	WindowHeight = 600
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
		WindowWidth,
		WindowHeight,
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

	ship, err := Load(r, "assets/ship.png")
	if err != nil {
		return err
	}

	missile, err := Load(r, "assets/sprite_1.png")
	if err != nil {
		return err
	}

	boss, err := Load(r, "assets/boss_ship.png")
	if err != nil {
		return err
	}

	frect := entities.NewFloatRectFactory().
		Create(ship, entities.Player, 375, 475, 50, 50)

	enemyPool, err := NewGameEnemyPoolFactory().CreateDefault([]*sdl.Texture{ship, boss})
	if err != nil {
		return err
	}

	var enemies []entities.Entity[float32, *sdl.FRect]
	for _, e := range enemyPool.Enemies() {
		enemies = append(enemies, e)
	}

	projectile := managers.NewProjectileManagerFactory().
		CreateInNegativeDirection(frect, enemies, missile)

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
			projectile.Spawn(r)
		},
		func() {
			projectile.Behave(r)
		},
	)

	if err := enemyPool.DrawEnemies(r); err != nil {
		return err
	}

	player := NewGamePlayer(
		frect,
		[]handlers.KeyboardEventHandler{projectileHandler},
		[]handlers.AfterEventHandler{movementHandler, projectileHandler},
	)
	if err := player.Draw(r); err != nil {
		return err
	}
	r.Present()

	if err := bg.Draw(r); err != nil {
		return err
	}

	loop := NewGameLoop(player, enemyPool, bg)
	return loop.Run(r)
}

func Load(r *sdl.Renderer, filename string) (*sdl.Texture, error) {
	img, err := img.Load(filename)
	if err != nil {
		return nil, err
	}
	return r.CreateTextureFromSurface(img)
}
