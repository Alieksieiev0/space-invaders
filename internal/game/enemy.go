package game

import (
	"math/rand"
	"time"

	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/Alieksieiev0/space-invaders/internal/managers"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	MinInRow    = 5
	MinInColumn = 3
)

type GameEnemy struct {
	entities.MoveableEntity[float32, *sdl.FRect]
	projectileManager managers.ProjectileManager
	isFired           bool
}

func (g *GameEnemy) RenderAfterEvent(r *sdl.Renderer) {
	if g.projectileManager == nil {
		return
	}
	if !g.isFired {
		go func() {
			g.isFired = true
			rand.NewSource(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			g.projectileManager.Spawn(r)
			g.isFired = false
		}()
	}
	g.projectileManager.Behave(r)
}

func (g *GameEnemy) AddTarget(entity entities.Entity[float32, *sdl.FRect]) {
	if g.projectileManager != nil {
		g.projectileManager.AddTarget(entity)
	}
}

func (g *GameEnemy) Fire(r *sdl.Renderer) {
	if g.projectileManager != nil {
		g.projectileManager.Spawn(r)
	}
}

func (g *GameEnemy) InitializeProjectileManager(texture *sdl.Texture) {
	g.projectileManager = managers.NewBidirectionalProjectileFactory().
		CreateInPositiveDirection(g, texture)
}

func NewGameEnemyFactory() GameEnemyFactory {
	return GameEnemyFactory{}
}

type GameEnemyFactory struct {
}

func (g GameEnemyFactory) Create(x, y, w, h float32) Enemy {
	return &GameEnemy{
		MoveableEntity: entities.NewFloatRectFactory().
			CreateWithoutTexture(entities.Enemy, x, y, w, h),
	}
}
