package main

import (
	"./notifications"
	"./world"
)

type solution struct {
	Noty *notifications.GodNotificator
}

func newSolution() (*solution, error) {
	noty, err := notifications.NewGodNotificator()
	if err != nil {
		return nil, err
	}
	return &solution{
		Noty: noty,
	}, nil
}

func (sol *solution) run() {
	w := world.NewWorld()
	// TODO
}
