package mechanics

type Collision interface {
	Collide(other Collision)
}
