package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputPath = "days/day2/input.txt"

var (
	opponentMarks = []string{
		"A", // Rock
		"B", // Paper
		"C", // Scissors
	}

	winMap = map[string]int{
		"Y": 0, // Paper
		"Z": 1, // Scissors
		"X": 2, // Rock
	}
	drawMap = map[string]int{
		"X": 0, // Rock
		"Y": 1, // Paper
		"Z": 2, // Scissors
	}

	markValues = map[string]int{
		"X": 1, // Rock
		"Y": 2, // Paper
		"Z": 3, // Scissors
	}
)

func calc(source io.Reader) (int, error) {
	score := 0
	r := bufio.NewReader(source)

	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, fmt.Errorf("reader read line: %w", err)
		}

		line := string(b)
		if line == "" {
			break
		}

		marks := strings.Split(line, " ")

		opponentMark, santaMark := marks[0], marks[1]
		score += markValues[santaMark]
		if opponentMarks[winMap[santaMark]] == opponentMark {
			score += 6
		} else if opponentMarks[drawMap[santaMark]] == opponentMark {
			score += 3
		}
	}

	return score, nil
}

func main() {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatalln(fmt.Errorf("os open: %w", err).Error())
	}

	result, err := calc(f)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("The result is: %d\n", result)
}
