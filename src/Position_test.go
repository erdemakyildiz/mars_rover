package main

import "testing"

func TestPosition_Add(t *testing.T) {
	position := Position{
		x: 3,
		y: 3,
	}
	position2 := Position{
		x: 1,
		y: 0,
	}

	result := position.Add(position2)

	if result.x != 4 && result.y != 3 {
		t.Error()
	}
}