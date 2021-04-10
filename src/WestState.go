package main

type WestState struct{}

func (s *WestState) Right(c *DirectionContext) {
	c.ChangeState(new(NorthState))
}

func (s *WestState) Left(c *DirectionContext) {
	c.ChangeState(new(SouthState))
}

func (s *WestState) GetPosition() Position {
	return Position{
		x: -1,
		y: 0,
	}
}

func (s *WestState) GetLabel() string {
	return "W"
}
