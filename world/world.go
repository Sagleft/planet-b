package world

// NewWorld - create new simulation
func NewWorld() *World {
	w := World{
		Name: "Planet B",
	}
	w.Transform.Position = Vector3{}.Zero()
	return &w
}
