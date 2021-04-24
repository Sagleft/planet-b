package world

type world struct {
	Name      string
	Transform transform
}

type transform struct {
	Position vector3
}

type vector3 struct {
	X, Y, Z float64
}

func (vector3) zero() vector3 {
	return vector3{0, 0, 0}
}
