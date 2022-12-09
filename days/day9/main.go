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

func solveP1(source io.Reader) int {
	s := bufio.NewScanner(source)
	r := newRope(1)

	for s.Scan() {
		lines := strings.Split(s.Text(), " ")
		dir, rawDelta := lines[0], lines[1]
		delta, err := strconv.Atoi(rawDelta)
		if err != nil {
			log.Fatal(err)
		}
		r.move(direction(dir), delta)
	}

	return r.getVisitedPositions()
}

func solveP2(source io.Reader) int {
	s := bufio.NewScanner(source)
	g := newRope(9)

	for s.Scan() {
		lines := strings.Split(s.Text(), " ")
		dir, rawDelta := lines[0], lines[1]
		delta, err := strconv.Atoi(rawDelta)
		if err != nil {
			log.Fatal(err)
		}
		g.move(direction(dir), delta)
	}

	return g.getVisitedPositions()
}

func main() {
	// part 1
	{
		f, err := os.Open("days/day9/input.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println("part 1:", solveP1(f))

		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}

	// part 2
	{
		f, err := os.Open("days/day9/input.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println("part 2:", solveP2(f))

		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
