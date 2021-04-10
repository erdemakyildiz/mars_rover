package main

import "strconv"

type Rover struct {
	plateau Plateau
	currentPosition Position
	direction DirectionContext
}

func (r *Rover) turnRight() {
	r.direction.Right()
}

func (r *Rover) turnLeft() {
	r.direction.Left()
}

func (r *Rover) moveForward() {
	newPosition := Move(r.currentPosition, r.direction.Position())

	if r.plateau.PositionIsValid(newPosition){
		r.currentPosition = newPosition
	}
}

func (r *Rover) getPosition() string {
	return strconv.Itoa(r.currentPosition.x) + " " + strconv.Itoa(r.currentPosition.y) + " " + r.direction.GetLabel()
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