package entities

import (
	"github.com/veandco/go-sdl2/sdl"
)

type RectSprite struct {
	rect       *sdl.Rect
	texture    *sdl.Texture
	stepX      int32
	stepY      int32
	entityType EntityType
}

func (r *RectSprite) Draw(renderer *sdl.Renderer) error {
	return renderer.Copy(r.texture, nil, r.rect)
}

func (r *RectSprite) StepX() int32 {
	return r.stepX
}

func (r *RectSprite) StepY() int32 {
	return r.stepY
}

func (r *RectSprite) Move(x int32, y int32) {
	r.rect.X += x
	r.rect.Y += y
}

func (r *RectSprite) Type() EntityType {
	return r.entityType
}

func (f *RectSprite) Dimensions() (x, y, w, h int32) {
	return f.rect.X, f.rect.Y, f.rect.W, f.rect.H
}

type FloatRectSprite struct {
	rect       *sdl.FRect
	texture    *sdl.Texture
	stepX      float32
	stepY      float32
	entityType EntityType
}

func (f *FloatRectSprite) Draw(renderer *sdl.Renderer) error {
	return renderer.CopyF(f.texture, nil, f.rect)
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

func (f *FloatRectSprite) Dimensions() (x, y, w, h float32) {
	return f.rect.X, f.rect.Y, f.rect.W, f.rect.H
}

func NewRectFactory() RectFactory {
	return RectFactory{}
}

type RectFactory struct {
}

func (r RectFactory) Create(
	texture *sdl.Texture,
	entityType EntityType,
	X, Y, W, H int32,
) MoveableEntity[int32] {
	return &RectSprite{
		rect:       &sdl.Rect{X: X, Y: Y, W: W, H: H},
		texture:    texture,
		stepX:      25,
		stepY:      25,
		entityType: entityType,
	}
}

func NewFloatRectFactory() FloatRectFactory {
	return FloatRectFactory{}
}

type FloatRectFactory struct {
}

func (r FloatRectFactory) Create(
	texture *sdl.Texture,
	entityType EntityType,
	X, Y, W, H float32,
) MoveableEntity[float32] {
	return &FloatRectSprite{
		rect:       &sdl.FRect{X: X, Y: Y, W: W, H: H},
		texture:    texture,
		stepX:      3,
		stepY:      3,
		entityType: entityType,
	}
}
