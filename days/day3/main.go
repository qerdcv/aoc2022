package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var ErrNoCollisionFound = errors.New("no collision found")

func checkInStash(val byte, bs []byte) bool {
	for _, b := range bs {
		if b == val {
			return true
		}
	}

	return false
}

func findCollision(a []byte, b []byte) (byte, bool) {
	for _, bCh := range b {
		if checkInStash(bCh, a) {
			return bCh, true
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

func calcGroupStashPriority(stashes [3]string) int {
	for _, ch := range stashes[0] {
		if strings.Contains(stashes[1], string(ch)) &&
			strings.Contains(stashes[2], string(ch)) {
			return collisionToPriority(byte(ch))
		}
	}

	return 0
}

func calcPriorityP2(source io.Reader) (int, error) {
	priority := 0
	groupSize := 3
	r := bufio.NewReader(source)
	iterating := true
	for iterating {
		groupStash := [3]string{
			"", "", "",
		}
		for i := 0; i < groupSize; i++ {
			b, _, err := r.ReadLine()
			if err != nil {
				if errors.Is(err, io.EOF) {
					iterating = false
				} else {
					return 0, fmt.Errorf("reader read line: %w", err)
				}
			}
			groupStash[i] = string(b)
		}
		priority += calcGroupStashPriority(groupStash)
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

	fmt.Println("Priority p1: ", priority)

	f.Seek(0, 0)
	priority, err = calcPriorityP2(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Priority p2: ", priority)
}
