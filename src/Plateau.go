package main

type Plateau struct {
	topRightX int
	topRightY int

	bottomLeftX int
	bottomLeftY int
}

func CratePlateau(topRightX, topRightY int) Plateau {
	return Plateau{
		topRightX:   topRightX,
		topRightY:   topRightY,
		bottomLeftX: 0,
		bottomLeftY: 0,
	}
}

func (p Plateau) PositionIsValid(position Position) bool {
	return position.x >= p.bottomLeftX &&
		position.x <= p.topRightX &&
		position.y >= p.bottomLeftY &&
		position.y <= p.topRightY
}
