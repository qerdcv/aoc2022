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

func solve(source io.Reader) int {
	s := bufio.NewScanner(source)
	c := newComputer()

	for s.Scan() {
		line := s.Text()
		sLine := strings.Split(line, " ")
		if len(sLine) == 1 {
			c.noop()
		} else {
			val, err := strconv.Atoi(sLine[1])
			if err != nil {
				log.Fatalln(err)
			}

			c.addX(val)
		}
	}

	return c.signalStrength
}

func main() {
	{
		f, err := os.Open("days/day10/input.txt")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("solution for p1: ", solve(f))
	}
}
