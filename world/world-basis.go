package world

// World - define your own universe
type World struct {
	Name      string
	Transform Transform
}

// Transform ..
type Transform struct {
	Position Vector3
}

// Vector3 - 3D vector
type Vector3 struct {
	X, Y, Z float64
}

// Zero - create point vector
func (Vector3) Zero() Vector3 {
	return Vector3{0, 0, 0}
}
