package main

import "testing"


func TestMove(t *testing.T) {
	position := Position{
		x: 2,
		y: 2,
	}

	north := DirectionContext{state: &NorthState{}}

	nextPosition := Move(north.Position(), position)

	if nextPosition.y != 3 && nextPosition.x != 2 {
		t.Error()
	}
}

func TestMove_2(t *testing.T) {
	position := Position{
		x: 1,
		y: 1,
	}
	west := DirectionContext{state: &WestState{}}

	nextPosition := Move(west.Position(), position)

	if nextPosition.y != 1 && nextPosition.x != 0 {
		t.Error()
	}
}