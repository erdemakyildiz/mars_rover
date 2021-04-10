package main

var North = Position{
	x: 0,
	y: 1,
}

var East = Position{
	x: 1,
	y: 0,
}

var South = Position{
	x: 0,
	y: -1,
}

var West = Position{
	x: -1,
	y: 0,
}

var directionList = [4]Position{North, East, South, West}
var flags = [4]string{"N", "E", "S", "W"}

func Rotate(r int, currentDirection string) string {
	for i := range flags {
		if flags[i] == currentDirection {
			var index = i + r

			if index < 0 {
				index = len(flags) - 1
			}

			if index > (len(flags) - 1) {
				index = 0
			}

			return flags[index]
		}
	}

	return currentDirection
}

func Move(currentDirection string, currentPosition Position) Position {
	return currentPosition.Add(directionList[index(flags, currentDirection)])
}

func index(slice [4]string, item string) int {
	for i, _ := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}