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

	sPointInRest := 0
	for { // until sPoint is not fallen into void
		sPoint := point{x: 500, y: 1}
		for { // until sPoint is not in rest
			if c.isPointTaken(point{sPoint.x, sPoint.y + 1}, false) {
				if !c.isPointTaken(point{sPoint.x - 1, sPoint.y + 1}, false) {
					sPoint = point{sPoint.x - 1, sPoint.y + 1}
					continue
				}

				if !c.isPointTaken(point{sPoint.x + 1, sPoint.y + 1}, false) {
					sPoint = point{sPoint.x + 1, sPoint.y + 1}
					continue
				}

				c.drawPoint(sPoint)
				sPointInRest += 1
				break
			}

			sPoint = point{sPoint.x, sPoint.y + 1}
			if sPoint.y > c.bottomLim {
				return sPointInRest
			}
		}
	}

}

func solveP2(source io.Reader) int {
	c := buildCave(source)
	sPointInRest := 0
	for { // while sPoint is not in start position
		sPoint := point{x: 500, y: 0}
		if c.isPointTaken(sPoint, true) {
			return sPointInRest
		}

		for { // until sPoint is not in rest state
			if c.isPointTaken(point{sPoint.x, sPoint.y + 1}, true) {
				if !c.isPointTaken(point{sPoint.x - 1, sPoint.y + 1}, true) {
					sPoint = point{sPoint.x - 1, sPoint.y + 1}
					continue
				}

				if !c.isPointTaken(point{sPoint.x + 1, sPoint.y + 1}, true) {
					sPoint = point{sPoint.x + 1, sPoint.y + 1}
					continue
				}

				c.drawPoint(sPoint)
				sPointInRest += 1
				break
			}

			sPoint = point{sPoint.x, sPoint.y + 1}
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
