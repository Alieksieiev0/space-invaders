package game

import (
	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/veandco/go-sdl2/sdl"
)

type GameEnemyPool struct {
	enemies []Enemy
}

func (g *GameEnemyPool) HandleAfterEvent() {
	for _, e := range g.enemies {
		e.HandleAfterEvent()
	}
}

func (g *GameEnemyPool) AddTarget(entity entities.Entity[float32, *sdl.FRect]) {
	for _, e := range g.enemies {
		e.AddTarget(entity)
	}
}

// in theory some of the targets might not be deleted yet(
func (g *GameEnemyPool) Targets() []entities.Entity[float32, *sdl.FRect] {
	if len(g.enemies) == 0 {
		return nil
	}
	return g.enemies[0].Targets()
}

func (g *GameEnemyPool) Enemies() []Enemy {
	return g.enemies
}

func (g *GameEnemyPool) DrawEnemies(r *sdl.Renderer) error {
	for _, e := range g.enemies {
		if e.IsAlive() {
			if err := e.Rotate(r, 180); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewGameEnemyPoolFactory() GameEnemyPoolFactory {
	return GameEnemyPoolFactory{}
}

type GameEnemyPoolFactory struct {
}

func (g GameEnemyPoolFactory) CreateDefault(textures []*sdl.Texture) (EnemyPool, error) {
	enemies, err := NewGameEnemyFactory().CreateMultipleWithRandomTexture(13, textures, 50, 50)
	if err != nil {
		return nil, err
	}

	return &GameEnemyPool{enemies: enemies}, nil
}
