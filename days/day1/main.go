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

func caloriesCalc(source io.Reader) (int, error) {
	result := 0
	r := bufio.NewReader(source)

main_loop:
	for {
		elfCal := 0
		for {
			b, _, err := r.ReadLine()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break main_loop
				}

				return 0, fmt.Errorf("reader read line: %w", err)
			}

			line := string(b)
			if line == "" {
				break
			}

			cal, err := strconv.Atoi(line)
			if err != nil {
				return 0, fmt.Errorf("strconv a to i: %w", err)
			}
			elfCal += cal
		}

		if elfCal > result {
			result = elfCal
		}
	}

	return result, nil
}

func main() {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatalln(fmt.Errorf("os open: %w", err).Error())
	}

	result, err := caloriesCalc(f)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("The result is: %d\n", result)
}
