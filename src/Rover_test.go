package main

import "testing"

func TestCommands(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{1,2},
		direction:       "N",
	}

	rover.commands("LMLMLMLMM")

	if rover.currentPosition.x != 1 &&
		rover.currentPosition.y != 3 &&
		rover.direction != "N" {

		t.Error()
	}
}

func TestCommands_2(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{3,3},
		direction:       "E",
	}

	rover.commands("MMRMMRMRRM")

	if rover.currentPosition.x != 5 &&
		rover.currentPosition.y != 1 &&
		rover.direction != "E" {

		t.Error()
	}
}

func TestCommands_3(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{3,3},
		direction:       "N",
	}

	rover.commands("MMMMMMMMMMR")

	if rover.currentPosition.x != 3 &&
		rover.currentPosition.y != 5 &&
		rover.direction != "E" {

		t.Error()
	}
}

func TestRotate(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{3,3},
		direction:       "E",
	}

	rover.turnLeft()

	if rover.direction != "N" {
		t.Error()
	}
}

func TestRotate_2(t *testing.T) {
	plateau := CratePlateau(5,5)

	rover := Rover{
		plateau:         plateau,
		currentPosition: Position{3,3},
		direction:       "E",
	}

	rover.turnRight()

	if rover.direction != "S" {
		t.Error()
	}
}