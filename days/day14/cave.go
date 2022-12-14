package main

import (
	"fmt"
)

type cave struct {
	c         map[string]bool
	bottomLim int
}

type point struct {
	x, y int
}

func newCave() cave {
	return cave{
		c: make(map[string]bool, 100),
	}
}

func (c *cave) drawLineToLeft(from point, dx int) {
	for x := from.x; x > from.x+dx-1; x-- {
		c.c[fmt.Sprintf("%d,%d", x, from.y)] = true
	}
}

func (c *cave) drawLineToRight(from point, dx int) {
	for x := from.x; x < from.x+dx+1; x++ {
		c.c[fmt.Sprintf("%d,%d", x, from.y)] = true
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
		c.c[fmt.Sprintf("%d,%d", from.x, y)] = true
	}
}

func (c *cave) drawLineToBot(from point, dy int) {
	for y := from.y; y < from.y+dy+1; y++ {
		if y > c.bottomLim {
			c.bottomLim = y
		}

		c.c[fmt.Sprintf("%d,%d", from.x, y)] = true
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
	c.c[fmt.Sprintf("%d,%d", p.x, p.y)] = true
}

func (c *cave) isPointTaken(p point, withBottomLim bool) bool {

	if withBottomLim {
		return c.c[fmt.Sprintf("%d,%d", p.x, p.y)] || p.y == c.bottomLim+2
	}

	return c.c[fmt.Sprintf("%d,%d", p.x, p.y)]
}
