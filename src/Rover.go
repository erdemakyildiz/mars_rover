package main

import "strconv"

type Rover struct {
	plateau Plateau
	currentPosition Position
	direction string
}

func (r *Rover) turnRight() {
	r.direction = Rotate(1, r.direction)
}

func (r *Rover) turnLeft() {
	r.direction = Rotate(-1, r.direction)
}

func (r *Rover) moveForward() {
	newPosition := Move(r.direction, r.currentPosition)

	if r.plateau.PositionIsValid(newPosition){
		r.currentPosition = newPosition
	}
}

func (r *Rover) getPosition() string {
	return strconv.Itoa(r.currentPosition.x) + " " + strconv.Itoa(r.currentPosition.y) + " " + r.direction
}

func (r *Rover) commands(cmd string) {
	for _, command := range cmd {
		c := string(command)

		switch c {
		case "R":
			r.turnRight()
		case "L":
			r.turnLeft()
		case "M":
			r.moveForward()
		}
	}
}