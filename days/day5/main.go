package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseRawStack(source io.Reader) ([][]byte, *bufio.Reader, error) {
	rawStacks := make([][]byte, 0)
	r := bufio.NewReader(source)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			return nil, nil, fmt.Errorf("reader readline: %w", err)
		}

		if string(b) == "" {
			break
		}

		rawStacks = append(rawStacks, b)
	}

	return rawStacks, r, nil
}
func parseStacks(source io.Reader) ([][]byte, *bufio.Reader, error) {
	rawStacks, r, err := parseRawStack(source)
	if err != nil {
		return nil, nil, fmt.Errorf("parse raw stack: %w", err)
	}

	rawStacksLen := len(rawStacks)
	stacks := make([][]byte, rawStacksLen)
	for i := rawStacksLen - 2; i >= 0; i-- {
		rawStack := rawStacks[i]
		for j := 0; j < len(stacks); j++ {
			offset := 4 * j
			limit := offset + 3
			if len(rawStack) < limit {
				break
			}

			item := rawStack[offset:limit]
			if string(item) == "   " {
				continue
			}
			stacks[j] = append(stacks[j], item[1])
		}
	}

	return stacks, r, nil
}

func executeCommands(stacks [][]byte, r *bufio.Reader) ([][]byte, error) {
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, fmt.Errorf("reader read line")
		}

		sCommand := strings.Split(string(b), " ")
		rawCnt, rawFrom, rawTo := sCommand[1], sCommand[3], sCommand[5]
		cnt, _ := strconv.Atoi(rawCnt)
		from, _ := strconv.Atoi(rawFrom)
		to, _ := strconv.Atoi(rawTo)

		// translate cnt to indexes
		from--
		to--

		for i := 0; i < cnt; i++ {
			fromStackLen := len(stacks[from])
			v := stacks[from][fromStackLen-1]
			stacks[from] = stacks[from][:fromStackLen-1]
			stacks[to] = append(stacks[to], v)
		}
	}

	return stacks, nil
}

func concatTops(stacks [][]byte) []byte {
	result := make([]byte, 0, len(stacks))

	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}

		result = append(result, stack[len(stack)-1])
	}

	return result
}

func calcP1(source io.Reader) ([]byte, error) {
	stacks, r, err := parseStacks(source)
	if err != nil {
		return nil, fmt.Errorf("parse stack: %w", err)
	}

	stacks, err = executeCommands(stacks, r)
	if err != nil {
		return nil, fmt.Errorf("parse stack: %w", err)
	}

	return concatTops(stacks), nil
}

func executeCommandsWithSavedDir(stacks [][]byte, r *bufio.Reader) ([][]byte, error) {
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, fmt.Errorf("reader read line")
		}

		sCommand := strings.Split(string(b), " ")
		rawCnt, rawFrom, rawTo := sCommand[1], sCommand[3], sCommand[5]
		cnt, _ := strconv.Atoi(rawCnt)
		from, _ := strconv.Atoi(rawFrom)
		to, _ := strconv.Atoi(rawTo)

		// translate cnt to indexes
		from--
		to--

		stacksFromLen := len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][stacksFromLen-cnt:]...)
		stacks[from] = stacks[from][:stacksFromLen-cnt]
	}

	return stacks, nil
}

func calcP2(source io.Reader) ([]byte, error) {
	stacks, r, err := parseStacks(source)
	if err != nil {
		return nil, fmt.Errorf("parse stack: %w", err)
	}

	stacks, err = executeCommandsWithSavedDir(stacks, r)
	if err != nil {
		return nil, fmt.Errorf("parse stack: %w", err)
	}

	return concatTops(stacks), nil
}

func main() {
	f, err := os.Open("days/day5/input.txt")
	if err != nil {
		panic(err)
	}

	result, err := calcP1(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1: ", string(result))

	f.Seek(0, 0)

	result, err = calcP2(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 2: ", string(result))
}
