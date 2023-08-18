package main

type direction int

const (
	n direction = iota
	s
	w
	e
)

const (
	maxx = 100
	maxy = 100
)

type Rover interface {
	Move(path string)
}

func NewRover(x, y int, direction direction) Rover {
	return &rover{
		x, y, direction,
	}
}

type rover struct {
	x int
	y int
	d direction
}

func (r *rover) moveCoord(x, y int) {

}

func (r *rover) Move(path string) {
	for _, d := range path {
		switch d {
		case 'f':
			switch r.d {
			case n:
				r.y++
			case s:
				r.y--
			case e:
				r.x++
			case w:
				r.x--
			default:
				panic("Invalid path")
			}
		case 'b':
			switch r.d {
			case s:
				r.y++
			case n:
				r.y--
			case w:
				r.x++
			case e:
				r.x--
			default:
				panic("Invalid path")
			}
		}
	}
}
