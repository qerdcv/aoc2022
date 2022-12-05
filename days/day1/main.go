package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const inputPath = "days/day1/input.txt"

func shiftTop(top [3]int, newTop int) [3]int {
	if newTop > top[2] {
		top[0], top[1], top[2] = top[1], top[2], newTop
		top[2] = newTop
	} else if newTop > top[1] {
		top[0], top[1] = top[1], newTop
	} else if newTop > top[0] {
		top[0] = newTop
	}
	return top
}

func caloriesCalc(source io.Reader) (int, [3]int, error) {
	top := [3]int{0, 0, 0}

	result := 0
	r := bufio.NewReader(source)

	iterating := true
	for iterating {
		elfCal := 0
		for {
			b, _, err := r.ReadLine()
			if err != nil {
				if errors.Is(err, io.EOF) {
					iterating = false
				} else {
					return 0, top, fmt.Errorf("reader read line: %w", err)
				}
			}

			line := string(b)
			if line == "" {
				break
			}

			cal, err := strconv.Atoi(line)
			if err != nil {
				return 0, top, fmt.Errorf("strconv a to i: %w", err)
			}
			elfCal += cal
		}

		top = shiftTop(top, elfCal)
		if elfCal > result {
			result = elfCal
		}
	}

	return result, top, nil
}

func main() {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatalln(fmt.Errorf("os open: %w", err).Error())
	}

	result, top, err := caloriesCalc(f)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("The result is: %d %d\n", result, top[0]+top[1]+top[2])
}
