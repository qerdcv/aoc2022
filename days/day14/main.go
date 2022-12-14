package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func newPointFromRaw(rawPoint string) (p point) {
	ps := strings.Split(rawPoint, ",")
	rawX := ps[0]
	rawY := ps[1]

	p.x, _ = strconv.Atoi(rawX)
	p.y, _ = strconv.Atoi(rawY)

	return p
}

func buildCave(source io.Reader) cave {
	c := newCave()

	s := bufio.NewScanner(source)
	for s.Scan() {
		line := s.Text()

		points := strings.Split(line, " -> ")
		prevPoint := newPointFromRaw(points[0])
		for _, p := range points[1:] {
			newPoint := newPointFromRaw(p)
			c.drawLine(prevPoint, newPoint)
			prevPoint = newPoint
		}
	}

	return c
}

func solveP1(source io.Reader) int {
	c := buildCave(source)

	sandInRest := 0
	for { // until sand is not fallen into void
		sand := point{x: 500, y: 1}
		for { // until sand is not in rest
			if c.isPointTaken(point{sand.x, sand.y + 1}, false) {
				if !c.isPointTaken(point{sand.x - 1, sand.y + 1}, false) {
					sand = point{sand.x - 1, sand.y + 1}
					continue
				}

				if !c.isPointTaken(point{sand.x + 1, sand.y + 1}, false) {
					sand = point{sand.x + 1, sand.y + 1}
					continue
				}

				c.drawPoint(sand)
				sandInRest += 1
				break
			}

			sand = point{sand.x, sand.y + 1}
			if sand.y > c.bottomLim {
				return sandInRest
			}
		}
	}
}

func solveP2(source io.Reader) int {
	c := buildCave(source)

	sandInRest := 0
	for { // while sand is not in start position
		sand := point{x: 500, y: 0}
		if c.isPointTaken(sand, true) {
			return sandInRest
		}

		for { // until sand is not in rest state
			if c.isPointTaken(point{sand.x, sand.y + 1}, true) {
				if !c.isPointTaken(point{sand.x - 1, sand.y + 1}, true) {
					sand = point{sand.x - 1, sand.y + 1}
					continue
				}

				if !c.isPointTaken(point{sand.x + 1, sand.y + 1}, true) {
					sand = point{sand.x + 1, sand.y + 1}
					continue
				}

				c.drawPoint(sand)
				sandInRest += 1
				break
			}

			sand = point{sand.x, sand.y + 1}
		}
	}
}

func main() {
	{
		f, err := os.Open("days/day14/input.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(solveP1(f))
	}
	{
		f, err := os.Open("days/day14/input.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(solveP2(f))
	}
}
