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

func NewProjectileManagerFactory() ProjectileManagerFactory {
	return ProjectileManagerFactory{}
}

type ProjectileManagerFactory struct {
}

func (p ProjectileManagerFactory) CreateInNegativeDirection(
	launcher entities.Entity[float32, *sdl.FRect],
	targets []entities.Entity[float32, *sdl.FRect],
	texture *sdl.Texture,
) EntityManager {
	return &ProjectileManager{
		launcher:  launcher,
		targets:   targets,
		texture:   texture,
		w:         width,
		h:         height,
		direction: negative,
	}
}

func (p ProjectileManagerFactory) CreateInPositiveDirection(
	launcher entities.Entity[float32, *sdl.FRect],
	targets []entities.Entity[float32, *sdl.FRect],
	texture *sdl.Texture,
) EntityManager {
	return &ProjectileManager{
		launcher:  launcher,
		targets:   targets,
		texture:   texture,
		w:         width,
		h:         height,
		direction: positive,
	}
}

type ProjectileManager struct {
	launcher    entities.Entity[float32, *sdl.FRect]
	texture     *sdl.Texture
	w           float32
	h           float32
	projectiles []entities.MoveableEntity[float32, *sdl.FRect]
	targets     []entities.Entity[float32, *sdl.FRect]
	direction   direction
}

func (p *ProjectileManager) Spawn(r *sdl.Renderer) error {
	x, y := p.coordinates()
	projectile := entities.NewFloatRectFactory().
		Create(p.texture, entities.Missile, x, y, p.w, p.h)
	if err := p.draw(r, projectile); err != nil {
		return err
	}
	p.projectiles = append(p.projectiles, projectile)
	return nil
}

func (p *ProjectileManager) coordinates() (float32, float32) {
	x, y, w, h := p.launcher.Dimensions()
	center := x + (w/2 - p.w/2)
	switch p.direction {
	case positive:
		return center, y + h + p.h
	case negative:
		return center, y - p.h
	}
	return 0, 0
}

func (p *ProjectileManager) draw(
	r *sdl.Renderer,
	pr entities.MoveableEntity[float32, *sdl.FRect],
) error {
	switch p.direction {
	case positive:
		return pr.Rotate(r, 180)
	case negative:
		return pr.Draw(r)
	}
	return nil
}

func (p *ProjectileManager) Behave(r *sdl.Renderer) error {
	for i, pr := range p.projectiles {
		switch p.direction {
		case positive:
			pr.Move(0, 1.00)
		case negative:
			pr.Move(0, -3.00)
		}
		if err := p.handleNewPosition(r, i, pr); err != nil {
			return err
		}
	}
	return nil
}

func (p *ProjectileManager) handleNewPosition(
	r *sdl.Renderer,
	i int,
	pr entities.MoveableEntity[float32, *sdl.FRect],
) error {
	for j, t := range p.targets {
		if pr.Intersect(t) {
			t.Destroy()
			p.projectiles = slices.Delete(p.projectiles, i, i+1)
			p.targets = slices.Delete(p.targets, j, j+1)
			return nil
		}
	}
	return pr.Draw(r)
}
