package entities

import (
	"github.com/veandco/go-sdl2/sdl"
)

type FloatRectSprite struct {
	rect       *sdl.FRect
	texture    *sdl.Texture
	stepX      float32
	stepY      float32
	entityType EntityType
	isAlive    bool
}

func (f *FloatRectSprite) Draw(renderer *sdl.Renderer) error {
	return renderer.CopyF(f.texture, nil, f.rect)
}

func (f *FloatRectSprite) Rotate(renderer *sdl.Renderer, angle float64) error {
	return renderer.CopyExF(f.texture, nil, f.rect, angle, nil, sdl.FLIP_NONE)
}

func (f *FloatRectSprite) Rect() *sdl.FRect {
	rectCopy := *f.rect
	return &rectCopy
}

func (f *FloatRectSprite) Intersect(entity Entity[float32, *sdl.FRect]) bool {
	return f.rect.HasIntersection(entity.Rect())
}

func (f *FloatRectSprite) StepX() float32 {
	return f.stepX
}

func (f *FloatRectSprite) StepY() float32 {
	return f.stepY
}

func (f *FloatRectSprite) Move(x float32, y float32) {
	f.rect.X += x
	f.rect.Y += y
}

func (f *FloatRectSprite) Type() EntityType {
	return f.entityType
}

func (f *FloatRectSprite) LoadTexture(texture *sdl.Texture) {
	f.texture = texture
}

func (f *FloatRectSprite) Dimensions() (x, y, w, h float32) {
	return f.rect.X, f.rect.Y, f.rect.W, f.rect.H
}

func (f *FloatRectSprite) Destroy() {
	f.isAlive = false
	f.rect = nil
}

func (f *FloatRectSprite) IsAlive() bool {
	return f.isAlive
}

func NewFloatRectFactory() FloatRectFactory {
	return FloatRectFactory{}
}

type FloatRectFactory struct {
}

func (r FloatRectFactory) Create(
	texture *sdl.Texture,
	entityType EntityType,
	x, y, w, h float32,
) MoveableEntity[float32, *sdl.FRect] {
	return &FloatRectSprite{
		rect:       &sdl.FRect{X: x, Y: y, W: w, H: h},
		texture:    texture,
		stepX:      3,
		stepY:      10,
		entityType: entityType,
		isAlive:    true,
	}
}

func (r FloatRectFactory) CreateWithoutTexture(
	entityType EntityType,
	x, y, w, h float32,
) MoveableEntity[float32, *sdl.FRect] {
	return &FloatRectSprite{
		rect:       &sdl.FRect{X: x, Y: y, W: w, H: h},
		stepX:      3,
		stepY:      10,
		entityType: entityType,
		isAlive:    true,
	}
}
