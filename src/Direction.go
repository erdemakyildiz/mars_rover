package main

type DirectionState interface {
	Right(c *DirectionContext)
	Left(c *DirectionContext)
	GetPosition() Position
	GetLabel() string
}

type DirectionContext struct {
	state DirectionState
}

func (c *DirectionContext) GetLabel() string {
	return c.state.GetLabel()
}

func (c *DirectionContext) Position() Position {
	return c.state.GetPosition()
}

func (c *DirectionContext) Right() {
	c.state.Right(c)
}

func (c *DirectionContext) Left() {
	c.state.Left(c)
}

func (c *DirectionContext) ChangeState(state DirectionState) {
	c.state = state
}

func Move(currentPosition Position, pos Position) Position {
	return currentPosition.Add(pos)
}

