package main

import "testing"

func TestCommands(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{1,2},
		direction:       DirectionContext{state: &NorthState{}},
	}

	rover.commands("LMLMLMLMM")

	if rover.currentPosition.x != 1 &&
		rover.currentPosition.y != 3 &&
		rover.direction.GetLabel() != "N" {

		t.Error()
	}
}

func TestCommands_2(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{3,3},
		direction:       DirectionContext{state: &EastState{}},
	}

	rover.commands("MMRMMRMRRM")

	if rover.currentPosition.x != 5 &&
		rover.currentPosition.y != 1 &&
		rover.direction.GetLabel() != "E" {

		t.Error()
	}
}

func TestCommands_3(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{3,3},
		direction:       DirectionContext{state: &NorthState{}},
	}

	rover.commands("MMMMMMMMMMR")

	if rover.currentPosition.x != 3 &&
		rover.currentPosition.y != 5 &&
		rover.direction.GetLabel() != "E" {

		t.Error()
	}
}

func TestRotate(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{3,3},
		direction:       DirectionContext{state: &EastState{}},
	}

	rover.turnLeft()

	if rover.direction.GetLabel() != "N" {
		t.Error()
	}
}

func TestRotate_2(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{3,3},
		direction:       DirectionContext{state: &EastState{}},
	}

	rover.turnRight()

	if rover.direction.GetLabel() != "S" {
		t.Error()
	}
}