package world

func newWorld() *world {
	w := world{
		Name: "Planet B",
	}
	w.Transform.Position = vector3{}.zero()
	return &w
}
