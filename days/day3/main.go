package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrNoCollisionFound = errors.New("no collision found")

func findCollision(a []byte, b []byte) (byte, bool) {
	for _, bCh := range b {
		for _, aCh := range a {
			if bCh == aCh {
				return bCh, true
			}
		}
	}

	return 0, false
}

func lowerCollisionToPriority(ascii int) int {
	// (a)97 = 1

	return ascii - 96
}

func upperCollisionToPriority(ascii int) int {
	// (A)65 = 27
	return ascii - 38
}

func collisionToPriority(c byte) int {
	ascii := int(c)
	if ascii >= 97 {
		return lowerCollisionToPriority(ascii)
	} else {
		return upperCollisionToPriority(ascii)
	}
}

func calcPriority(source io.Reader) (int, error) {
	priority := 0
	r := bufio.NewReader(source)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, fmt.Errorf("reader read line: %w", err)
		}
		size := len(b) / 2
		c1, c2 := b[:size], b[size:]

		collision, ok := findCollision(c1, c2)
		if !ok {
			return 0, ErrNoCollisionFound
		}
		priority += collisionToPriority(collision)
	}

	return priority, nil
}

func main() {
	f, err := os.Open("days/day3/input.txt")
	if err != nil {
		panic(err)
	}

	priority, err := calcPriority(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(priority)
}
