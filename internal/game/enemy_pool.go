package game

import (
	"math"

	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/veandco/go-sdl2/sdl"
)

type enemyState int

const (
	standing enemyState = iota
	movingLeft
	movingRight
	movingBottom
)

type GameEnemyPool struct {
	enemies [][]Enemy
	state   enemyState
}

func (g *GameEnemyPool) HandleAfterEvent() {
	g.clearDead()
	switch g.state {
	case standing:
		g.state = movingRight
	case movingRight:
		if g.reachedRightBorder() {
			g.state = movingBottom
			return
		}
		for _, row := range g.enemies {
			for _, e := range row {
				e.Move(e.StepX()/2, 0)
			}
		}
	case movingLeft:
		if g.reachedLeftBorder() {
			g.state = movingBottom
			return
		}
		for _, row := range g.enemies {
			for _, e := range row {
				e.Move(-e.StepX()/2, 0)
			}
		}
	case movingBottom:
		for _, row := range g.enemies {
			for _, e := range row {
				e.Move(0, e.StepY())
			}
		}
		if g.reachedRightBorder() {
			g.state = movingLeft
			return
		}
		g.state = movingRight
	}
}

func (g *GameEnemyPool) clearDead() {
	for i, row := range g.enemies {
		j := 0
		for _, e := range row {
			if e.IsAlive() {
				row[j] = e
				j++
			}
		}
		for el := j; el < len(row); el++ {
			row[el] = nil
		}
		g.enemies[i] = row[:j]
	}
}

func (g *GameEnemyPool) reachedRightBorder() bool {
	x, _, w, _ := g.lastDimensions()
	return x+w >= WindowWidth
}

func (g *GameEnemyPool) reachedLeftBorder() bool {
	x, _, _, _ := g.firstDimensions()
	return x <= 0
}

func (g *GameEnemyPool) firstDimensions() (float32, float32, float32, float32) {
	var first Enemy
	lowestX := math.MaxFloat32
	for _, row := range g.enemies {
		if len(row) == 0 {
			continue
		}

		x, _, _, _ := row[0].Dimensions()
		if x < float32(lowestX) {
			lowestX = float64(x)
			first = row[0]
		}
	}
	return first.Dimensions()
}

func (g *GameEnemyPool) lastDimensions() (float32, float32, float32, float32) {
	var last Enemy
	var highestX float32
	for _, row := range g.enemies {
		if len(row) == 0 {
			continue
		}
		x, _, _, _ := row[len(row)-1].Dimensions()
		if x > highestX {
			highestX = x
			last = row[len(row)-1]
		}
	}
	return last.Dimensions()
}

func (g *GameEnemyPool) RenderAfterEvent(r *sdl.Renderer) {
	for _, row := range g.enemies {
		for _, e := range row {
			e.RenderAfterEvent(r)
		}
	}
}

func (g *GameEnemyPool) Enemies() [][]Enemy {
	return g.enemies
}

func (g *GameEnemyPool) Shooters() []Enemy {
	return g.enemies[0]
}

func (g *GameEnemyPool) DrawEnemies(r *sdl.Renderer) error {
	for _, row := range g.enemies {
		for _, e := range row {
			if e.IsAlive() {
				if err := e.Rotate(r, 180); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (g *GameEnemyPool) AddTarget(entity entities.Entity[float32, *sdl.FRect]) {
	for _, e := range g.Shooters() {
		e.AddTarget(entity)
	}
}

func (g *GameEnemyPool) Fire(r *sdl.Renderer) {
	for _, e := range g.Shooters() {
		e.Fire(r)
	}
}

func (g *GameEnemyPool) InitializeProjectileManager(texture *sdl.Texture) {
	for _, e := range g.Shooters() {
		e.InitializeProjectileManager(texture)
	}
}

func NewGameEnemyPoolFactory() GameEnemyPoolFactory {
	return GameEnemyPoolFactory{}
}

type GameEnemyPoolFactory struct {
}

func (g GameEnemyPoolFactory) CreateDefault(shooter *sdl.Texture, regular *sdl.Texture) EnemyPool {
	enemies := NewDefaultEnemyLayout().Generate()
	for i, row := range enemies {
		for _, e := range row {
			if i == 0 {
				e.LoadTexture(shooter)
				continue
			}
			e.LoadTexture(regular)
		}
	}

	return &GameEnemyPool{enemies: enemies, state: standing}
}
