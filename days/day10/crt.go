package main

import "fmt"

const (
	screenWidth  = 39
	screenHeight = 5
)

type crt struct {
	x, y int
}

func (c *crt) draw(pixelPosition int) {
	if pixelPosition >= c.x-1 && pixelPosition <= c.x+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	c.x++

	if c.x > screenWidth {
		fmt.Print("\n")
		c.y++
		c.x = 0
	}

	if c.y > screenHeight {
		fmt.Print("\n")
		c.y = 0
	}
}
