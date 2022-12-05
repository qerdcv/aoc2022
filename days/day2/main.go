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
	markValues = map[string]int{
		"X": 1, // Rock
		"Y": 2, // Paper
		"Z": 3, // Scissors
	}
)

func calc(source io.Reader) (int, error) {
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
	)

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

		marks := strings.Split(string(b), " ")
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

func calcP2(source io.Reader) (int, error) {
	// X LOSE
	// Y Draw
	// Z WIN
	var (
		winMap = map[string]string{
			"A": "Y", // Paper
			"B": "Z", // Scissors
			"C": "X", // Rock
		}
		loseMap = map[string]string{
			"A": "Z", // Rock
			"B": "X", // Paper
			"C": "Y", // Scissors
		}
		drawMap = map[string]string{
			"A": "X", // Rock
			"B": "Y", // Paper
			"C": "Z", // Scissors
		}
	)

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

		marks := strings.Split(string(b), " ")
		opponentMark, santaMark := marks[0], marks[1]
		switch santaMark {
		case "X": // lose
			score += markValues[loseMap[opponentMark]]
		case "Y": // draw
			score += markValues[drawMap[opponentMark]] + 3
		case "Z": // win
			score += markValues[winMap[opponentMark]] + 6
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

	log.Printf("The result p.1 is: %d\n", result)

	f.Seek(0, 0)
	result, err = calcP2(f)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("The result p.2 is: %d\n", result)
}
