package game

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
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

	background, err := Load(r, "assets/background.png")
	if err != nil {
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

	bg := entities.NewBackgroundFactory().Create(background)

	enemyPool := NewGameEnemyPoolFactory().CreateDefault(boss, ship)
	enemyPool.InitializeProjectileManager(missile)
	player := setupPlayer(r, ship, missile, enemyPool)
	enemyPool.AddTarget(player)

	if err := bg.Draw(r); err != nil {
		return err
	}
	if err := enemyPool.DrawEnemies(r); err != nil {
		return err
	}
	if err := player.Draw(r); err != nil {
		return err
	}
	r.Present()

	loop := NewGameLoop(player, enemyPool, bg)
	return loop.Run(r)
}

func setupPlayer(r *sdl.Renderer, ship *sdl.Texture, missile *sdl.Texture, pool EnemyPool) Player {
	player := NewGamePlayer(
		entities.NewFloatRectFactory().Create(ship, entities.Player, 375, 475, 55, 55),
	)
	player.InitializeProjectileManager(missile)
	for _, row := range pool.Enemies() {
		for _, e := range row {
			player.AddTarget(e)
		}
	}
	return player
}

func Load(r *sdl.Renderer, filename string) (*sdl.Texture, error) {
	img, err := img.Load(filename)
	if err != nil {
		return nil, err
	}
	return r.CreateTextureFromSurface(img)
}
