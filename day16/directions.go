package day16

var (
	UP    = Direction{-1, 0}
	LEFT  = Direction{0, -1}
	DOWN  = Direction{1, 0}
	RIGHT = Direction{0, 1}
)

var TURN_LEFT = map[Direction]Direction{
	UP:    LEFT,
	LEFT:  DOWN,
	DOWN:  RIGHT,
	RIGHT: UP,
}

var TURN_RIGHT = map[Direction]Direction{
	UP:    RIGHT,
	LEFT:  UP,
	DOWN:  LEFT,
	RIGHT: DOWN,
}

var DIRECTIONS = []Direction{UP, LEFT, DOWN, RIGHT}
