package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func checkUniq(b []byte) bool {
	fmt.Println(string(b))
	for idx, ch1 := range b {
		for _, ch2 := range b[idx+1:] {
			if ch1 == ch2 {
				return false
			}
		}
	}
	return true
}

func findMarker(source io.ReadSeeker) (int, error) {
	buf := make([]byte, 4)
	seeker := 0
	for {
		n, err := source.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, fmt.Errorf("source read: %w", err)
		}

		if n < 4 {
			break
		}

		if checkUniq(buf) {
			return seeker + 4, nil
		}

		seeker++
		if _, err = source.Seek(int64(seeker), 0); err != nil {
			return 0, fmt.Errorf("source seek: %w", err)
		}
	}

	return seeker, nil
}

func findMessage(source io.ReadSeeker) (int, error) {
	buf := make([]byte, 14)
	seeker := 0
	for {
		n, err := source.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, fmt.Errorf("source read: %w", err)
		}

		if n < 14 {
			break
		}

		if checkUniq(buf) {
			return seeker + 14, nil
		}

		seeker++
		if _, err = source.Seek(int64(seeker), 0); err != nil {
			return 0, fmt.Errorf("source seek: %w", err)
		}
	}

	return seeker, nil
}

func main() {
	f, err := os.Open("days/day6/input.txt")
	if err != nil {
		panic(err)
	}

	marker, err := findMarker(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start marker: ", marker)

	f.Seek(0, 0)
	messageMarker, err := findMessage(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("message marker: ", messageMarker)
}
