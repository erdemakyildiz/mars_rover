package main

type EastState struct{}

func (s *EastState) Right(c *DirectionContext) {
	c.ChangeState(new(SouthState))
}

func (s *EastState) Left(c *DirectionContext) {
	c.ChangeState(new(NorthState))
}

func (s *EastState) GetPosition() Position {
	return Position{
		x: 1,
		y: 0,
	}
}

func (s *EastState) GetLabel() string {
	return "E"
}
