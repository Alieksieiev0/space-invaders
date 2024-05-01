package game

import (
	"fmt"
	"math/rand"

	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/Alieksieiev0/space-invaders/internal/handlers"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	MinInRow         = 5
	MinInColumn      = 3
	SpaceCoefficient = 3
	WidthPadding     = 100
	HeightPadding    = 50
)

type GameEnemy struct {
	entities.MoveableEntity[float32, *sdl.FRect]
	targets            []entities.Entity[float32, *sdl.FRect]
	afterEventHandlers []handlers.AfterEventHandler
}

func (b *GameEnemy) HandleAfterEvent() {
	for _, h := range b.afterEventHandlers {
		h.HandleAfterEvent()
	}
}

func (g *GameEnemy) AddTarget(entity entities.Entity[float32, *sdl.FRect]) {
	g.targets = append(g.targets, entity)
}

func (g *GameEnemy) Targets() []entities.Entity[float32, *sdl.FRect] {
	return g.targets
}

func NewGameEnemyFactory() GameEnemyFactory {
	return GameEnemyFactory{}
}

type GameEnemyFactory struct {
}

func (g GameEnemyFactory) Create(texture *sdl.Texture, x, y, w, h float32) Enemy {
	return &GameEnemy{
		MoveableEntity: entities.NewFloatRectFactory().Create(texture, entities.Enemy, x, y, w, h),
	}
}

func (g GameEnemyFactory) CreateMultipleWithRandomTexture(
	number int,
	textures []*sdl.Texture,
	w, h float32,
) ([]Enemy, error) {
	ws, hs := g.spaces(w, h)
	rowSize, _, err := g.determineLayout(w, h, ws, hs)
	if err != nil {
		return nil, err
	}
	columnsNumber := (number / rowSize) + 1
	var enemies []Enemy

	lastY := float32(HeightPadding)
	for i := 0; i < columnsNumber; i++ {
		lastX := float32(WidthPadding)
		for j := 0; j < rowSize; j++ {
			if len(enemies) >= number {
				break
			}
			e := g.Create(textures[rand.Intn(len(textures))], lastX, lastY, w, h)
			lastX += w + ws
			enemies = append(enemies, e)
		}
		lastY += h + hs
	}

	return enemies, nil
}

func (g GameEnemyFactory) spaces(w, h float32) (float32, float32) {
	return w / SpaceCoefficient, h / SpaceCoefficient
}

func (g GameEnemyFactory) determineLayout(w, h, ws, hs float32) (int, int, error) {
	rowSize := (WindowWidth - WidthPadding*2) / (w + ws)
	columnSize := (WindowHeight - HeightPadding*2) / (h + hs)
	if rowSize < MinInRow || columnSize < MinInColumn {
		return 0, 0, fmt.Errorf("size for enemy sprite is too big")
	}
	return int(rowSize), int(columnSize), nil
}
