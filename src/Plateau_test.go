package main

import "testing"

func TestCratePlateau(t *testing.T) {
	plateau := CratePlateau(5, 10)

	if plateau.bottomLeftX != 0 &&
		plateau.bottomLeftY != 0 &&
		plateau.topRightX != 5 &&
		plateau.topRightY != 10 {

		t.Error()
	}
}

func TestPlateau_PositionIsValid(t *testing.T) {
	plateau := CratePlateau(5, 5)
	newPosition := Position{
		x: 6,
		y: 6,
	}

	if plateau.PositionIsValid(newPosition){
		t.Error()
	}
}

func TestPlateau_PositionIsValid_2(t *testing.T) {
	plateau := CratePlateau(5, 5)
	newPosition := Position{
		x: 3,
		y: 3,
	}

	if !plateau.PositionIsValid(newPosition){
		t.Error()
	}
}