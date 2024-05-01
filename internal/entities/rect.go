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
	isAlive    bool
}

func (r *RectSprite) Draw(renderer *sdl.Renderer) error {
	return renderer.Copy(r.texture, nil, r.rect)
}

func (r *RectSprite) Rotate(renderer *sdl.Renderer, angle float64) error {
	return renderer.CopyEx(r.texture, nil, r.rect, angle, nil, sdl.FLIP_NONE)
}

func (r *RectSprite) Rect() *sdl.Rect {
	rectCopy := *r.rect
	return &rectCopy
}

func (r *RectSprite) Intersect(entity Entity[int32, *sdl.Rect]) bool {
	return r.rect.HasIntersection(entity.Rect())
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

func (r *RectSprite) Dimensions() (x, y, w, h int32) {
	return r.rect.X, r.rect.Y, r.rect.W, r.rect.H
}

func (r *RectSprite) Destroy() {
	r.isAlive = false
	r.rect = nil
}

func (r *RectSprite) IsAlive() bool {
	return r.isAlive
}

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

func NewRectFactory() RectFactory {
	return RectFactory{}
}

type RectFactory struct {
}

func (r RectFactory) Create(
	texture *sdl.Texture,
	entityType EntityType,
	x, y, w, h int32,
) MoveableEntity[int32, *sdl.Rect] {
	return &RectSprite{
		rect:       &sdl.Rect{X: x, Y: y, W: w, H: h},
		texture:    texture,
		stepX:      3,
		stepY:      3,
		entityType: entityType,
		isAlive:    true,
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
	x, y, w, h float32,
) MoveableEntity[float32, *sdl.FRect] {
	return &FloatRectSprite{
		rect:       &sdl.FRect{X: x, Y: y, W: w, H: h},
		texture:    texture,
		stepX:      3,
		stepY:      3,
		entityType: entityType,
		isAlive:    true,
	}
}
