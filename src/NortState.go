package main

type NorthState struct{}

func (s *NorthState) Right(c *DirectionContext) {
	c.ChangeState(new(EastState))
}

func (s *NorthState) Left(c *DirectionContext) {
	c.ChangeState(new(WestState))
}

func (s *NorthState) GetPosition() Position {
	return Position{
		x: 0,
		y: 1,
	}
}

func (s *NorthState) GetLabel() string {
	return "N"
}
