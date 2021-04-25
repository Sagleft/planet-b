package main

import (
	"./notifications"
	"./world"
)

type solution struct {
	Noty notifications.GodNotificator
}

func newSolution() solution {
	return solution{}
}

func (sol *solution) run() {
	w := world.NewWorld()
	// TODO
}
