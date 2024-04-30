package mechanics

type Movement[T int32 | float32] interface {
	StepX() T
	StepY() T
	Move(x T, y T)
}

type Collision interface {
	Collide(other Collision)
}
