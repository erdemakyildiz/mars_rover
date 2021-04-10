package main

import "testing"

func TestRotate_Right(t *testing.T) {
	result := Rotate(1, "N")

	if result != "E" {
		t.Error()
	}
}

func TestRotate_Left(t *testing.T) {
	result := Rotate(-1, "N")

	if result != "W" {
		t.Error()
	}
}

func TestRotate_Left_2(t *testing.T) {
	result := Rotate(-1, "W")

	if result != "S" {
		t.Error()
	}
}

func TestMove(t *testing.T) {
	position := Position{
		x: 2,
		y: 2,
	}

	nextPosition := Move("N", position)

	if nextPosition.y != 3 && nextPosition.x != 2 {
		t.Error()
	}
}

func TestMove_2(t *testing.T) {
	position := Position{
		x: 1,
		y: 1,
	}

	nextPosition := Move("W", position)

	if nextPosition.y != 1 && nextPosition.x != 0 {
		t.Error()
	}
}