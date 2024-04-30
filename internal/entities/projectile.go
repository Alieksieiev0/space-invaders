package entities

import (
	"github.com/veandco/go-sdl2/sdl"
)

func NewBidirectionalProjectile(
	launcher Entity[float32],
	texture *sdl.Texture,
	width, height float32,
) Projectile {
	return &BidirectionalProjectile{
		launcher: launcher,
		texture:  texture,
		width:    width,
		height:   height,
	}
}

type BidirectionalProjectile struct {
	launcher    Entity[float32]
	texture     *sdl.Texture
	width       float32
	height      float32
	projectiles []MoveableEntity[float32]
}

func (b *BidirectionalProjectile) Fire(r *sdl.Renderer) {
	x, y, width, _ := b.launcher.Dimensions()
	projectile := NewFloatRectFactory().Create(b.texture, Missile, x+(width/2-b.width/2), y-b.height, b.width, b.height)
	projectile.Draw(r)
	b.projectiles = append(b.projectiles, projectile)
}

func (b *BidirectionalProjectile) Move(r *sdl.Renderer) {
	for _, p := range b.projectiles {
		p.Move(0, -1.00)
		p.Draw(r)
	}
}
