package main

import (
	"fmt"
)

type item int

const (
	stone item = iota + 1
	sand
)

type cave struct {
	c                            map[string]item
	bottomLim, leftLim, rightLim int
}

type point struct {
	x, y int
}

func newCave() cave {
	return cave{
		c: make(map[string]item, 100),
	}
}

func (c *cave) drawLineToLeft(from point, dx int) {
	for x := from.x; x > from.x+dx-1; x-- {
		if x > c.rightLim {
			c.rightLim = x
		}

		c.c[fmt.Sprintf("%d,%d", x, from.y)] = stone
	}
}

func (c *cave) drawLineToRight(from point, dx int) {
	for x := from.x; x < from.x+dx+1; x++ {
		if x < c.leftLim || c.leftLim == 0 {
			c.leftLim = x
		}

		c.c[fmt.Sprintf("%d,%d", x, from.y)] = stone
	}
}

func (c *cave) drawHorizontalLine(from point, dx int) {
	if dx > 0 {
		c.drawLineToRight(from, dx)
		return
	}

	c.drawLineToLeft(from, dx)
}

func (c *cave) drawLineToTop(from point, dy int) {
	for y := from.y; y > from.y+dy-1; y-- {
		c.c[fmt.Sprintf("%d,%d", from.x, y)] = stone
	}
}

func (c *cave) drawLineToBot(from point, dy int) {
	for y := from.y; y < from.y+dy+1; y++ {
		if y > c.bottomLim {
			c.bottomLim = y
		}

		c.c[fmt.Sprintf("%d,%d", from.x, y)] = stone
	}
}

func (c *cave) drawVerticalLine(from point, dy int) {
	if dy > 0 {
		c.drawLineToBot(from, dy)
		return
	}

	c.drawLineToTop(from, dy)
}

func (c *cave) drawLine(from, to point) {
	dx, dy := to.x-from.x, to.y-from.y

	if dx != 0 {
		c.drawHorizontalLine(from, dx)
		return
	}

	if dy != 0 {
		c.drawVerticalLine(from, dy)
		return
	}
}

func (c *cave) drawPoint(p point) {
	c.c[fmt.Sprintf("%d,%d", p.x, p.y)] = sand
}

func (c *cave) isPointTaken(p point, withBottomLim bool) bool {
	i := c.c[fmt.Sprintf("%d,%d", p.x, p.y)]
	isTaken := i == stone || i == sand

	if withBottomLim {
		return isTaken || p.y == c.bottomLim+2
	}

	return isTaken
}

func (c *cave) print() {
	for y := 0; y < c.bottomLim+3; y++ {
		for x := c.leftLim - 200; x < c.rightLim+200; x++ {
			switch c.c[fmt.Sprintf("%d,%d", x, y)] {
			case sand:
				fmt.Print("o")
			case stone:
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}
}
