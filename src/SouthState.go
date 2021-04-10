package main

type SouthState struct{}

func (s *SouthState) Right(c *DirectionContext) {
	c.ChangeState(new(WestState))
}

func (s *SouthState) Left(c *DirectionContext) {
	c.ChangeState(new(EastState))
}

func (s *SouthState) GetPosition() Position {
	return Position{
		x: 0,
		y: -1,
	}
}

func (s *SouthState) GetLabel() string {
	return "S"
}
