package managers

import (
	"slices"

	"github.com/Alieksieiev0/space-invaders/internal/entities"
	"github.com/veandco/go-sdl2/sdl"
)

type direction int64

const (
	positive direction = iota
	negative
)

const (
	width  = 25
	height = 25
)

func NewBidirectionalProjectileFactory() BidirectionalProjectileFactory {
	return BidirectionalProjectileFactory{}
}

type BidirectionalProjectileFactory struct {
}

func (b BidirectionalProjectileFactory) create(
	launcher entities.Entity[float32, *sdl.FRect],
	texture *sdl.Texture,
) *BidirectionalProjectile {
	return &BidirectionalProjectile{
		launcher: launcher,
		texture:  texture,
		w:        width,
		h:        height,
	}
}

func (b BidirectionalProjectileFactory) CreateInNegativeDirection(
	launcher entities.Entity[float32, *sdl.FRect],
	texture *sdl.Texture,
) ProjectileManager {
	p := b.create(launcher, texture)
	p.direction = negative

	return p
}

func (b BidirectionalProjectileFactory) CreateInPositiveDirection(
	launcher entities.Entity[float32, *sdl.FRect],
	texture *sdl.Texture,
) ProjectileManager {
	p := b.create(launcher, texture)
	p.direction = positive

	return p
}

type BidirectionalProjectile struct {
	launcher   entities.Entity[float32, *sdl.FRect]
	texture    *sdl.Texture
	w          float32
	h          float32
	projectile entities.MoveableEntity[float32, *sdl.FRect]
	targets    []entities.Entity[float32, *sdl.FRect]
	direction  direction
}

func (b *BidirectionalProjectile) AddTarget(entity entities.Entity[float32, *sdl.FRect]) {
	b.targets = append(b.targets, entity)
}

func (b *BidirectionalProjectile) Spawn(r *sdl.Renderer) error {
	if b.projectile != nil || !b.launcher.IsAlive() {
		return nil
	}
	x, y := b.coordinates()
	projectile := entities.NewFloatRectFactory().
		Create(b.texture, entities.Missile, x, y, b.w, b.h)
	if err := b.draw(r, projectile); err != nil {
		return err
	}
	b.projectile = projectile
	return nil
}

func (b *BidirectionalProjectile) coordinates() (float32, float32) {
	x, y, w, _ := b.launcher.Dimensions()
	center := x + (w/2 - b.w/2)
	switch b.direction {
	case positive:
		return center, y + b.h
	case negative:
		return center, y - b.h
	}
	return 0, 0
}

func (b *BidirectionalProjectile) draw(
	r *sdl.Renderer,
	pr entities.MoveableEntity[float32, *sdl.FRect],
) error {
	switch b.direction {
	case positive:
		return pr.Rotate(r, 180)
	case negative:
		return pr.Draw(r)
	}
	return nil
}

func (b *BidirectionalProjectile) Behave(r *sdl.Renderer) error {
	if b.projectile == nil {
		return nil
	}
	switch b.direction {
	case positive:
		b.projectile.Move(0, 2.00)
	case negative:
		b.projectile.Move(0, -2.00)
	}
	if err := b.handleNewPosition(r); err != nil {
		return err
	}
	return nil
}

func (b *BidirectionalProjectile) handleNewPosition(r *sdl.Renderer) error {
	for j, t := range b.targets {
		if t == nil || !t.IsAlive() {
			b.targets = slices.Delete(b.targets, j, j+1)
			break
		}
		if b.projectile.Intersect(t) {
			t.Destroy()
			b.projectile = nil
			b.targets = slices.Delete(b.targets, j, j+1)
			return nil
		}
	}
	rect := b.projectile.Rect()
	if rect.Y <= 0 || rect.Y >= 600 || rect.X <= 0 || rect.X >= 800 {
		b.projectile = nil
		return nil
	}
	return b.draw(r, b.projectile)
}
