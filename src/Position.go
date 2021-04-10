package main

type Position struct {
	x, y int
}

func (p *Position) Add(position Position) Position{
	return Position{
		x: p.x + position.x,
		y: p.y + position.y,
	}
}