package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	updateSize  uint = 30_000_000
	totalFsSize uint = 70_000_000
)

func buildFs(source io.Reader) (*fs, error) {
	s := bufio.NewScanner(source)

	root := newFs()
	currentFs := root
	for s.Scan() {
		line := s.Text()

		cmd := strings.Split(line, " ")
		switch cmd[0] {
		case "":
			return root, nil
		case "$":
			switch cmd[1] {
			case "cd":
				currentFs = currentFs.content[cmd[2]]
			default:
				continue
			}
		case "dir":
			currentFs.addDir(cmd[1])
		default:
			size, err := strconv.Atoi(cmd[0])
			if err != nil {
				return nil, fmt.Errorf("strconv atoi: %w", err)
			}
			currentFs.addFile(cmd[1], uint(size))
		}
	}

	return root, nil
}

func iterFsTreeP1(f *fs) uint {
	var sizeSum uint = 0
	for name, innerFs := range f.content {
		if isNameReserved(name) {
			continue
		}

		if innerFs.t == dir {
			if innerFs.size <= 100000 {
				sizeSum += innerFs.size
			}

			sizeSum += iterFsTreeP1(innerFs)
		}
	}

	return sizeSum
}

func iterFsTreeP2(f *fs, requiredSpace uint, delPretSize uint) uint {
	for name, innerFs := range f.content {
		if isNameReserved(name) {
			continue
		}

		if innerFs.t == dir {
			if size := innerFs.size; size < delPretSize && size >= requiredSpace {
				delPretSize = size
			}

			if newDelPretSize := iterFsTreeP2(innerFs, requiredSpace, delPretSize); newDelPretSize < delPretSize {
				delPretSize = newDelPretSize
			}
		}
	}

	return delPretSize
}

func main() {
	f, err := os.Open("days/day7/input.txt")
	if err != nil {
		panic(err)
	}

	fileSystem, err := buildFs(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("part 1 -", iterFsTreeP1(fileSystem))
	fmt.Println("part 2 -", iterFsTreeP2(fileSystem, fileSystem.size-(totalFsSize-updateSize), totalFsSize))
}
