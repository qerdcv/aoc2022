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

func buildFs(source io.Reader) (*fs, error) {
	r := bufio.NewReader(source)

	root := newFs()
	currentFs := root
	scanning := true
	for scanning {
		b, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				scanning = false
			} else {
				return nil, fmt.Errorf("read all: %w", err)
			}
		}

		line := string(b)
		cmd := strings.Split(line, " ")
		switch cmd[0] {
		case "":
			scanning = false
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

// TODO: make it without global variable
var sizeSum uint

func iterFsTreeP1(f *fs) {
	for name, innerFs := range f.content {
		if name == "/" || name == ".." || name == "." {
			continue
		}

		if innerFs.t == elTypeDir {
			if innerFs.size <= 100000 {
				sizeSum += innerFs.size
			}

			iterFsTreeP1(innerFs)
		}
	}
}

const (
	updateSize  uint = 30_000_000
	totalFsSize uint = 70_000_000
)

// TODO: make it without global variable
var currentDeletePretendent uint = totalFsSize

func iterFsTreeP2(f *fs, requiredSpace uint) {
	for name, innerFs := range f.content {
		if name == "/" || name == ".." || name == "." {
			continue
		}

		if innerFs.t == elTypeDir {
			if size := innerFs.size; size < currentDeletePretendent && size >= requiredSpace {
				currentDeletePretendent = size
			}
			iterFsTreeP2(innerFs, requiredSpace)
		}
	}
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

	iterFsTreeP1(fileSystem)
	fmt.Println("part 1 -", sizeSum)

	iterFsTreeP2(fileSystem, fileSystem.size-(totalFsSize-updateSize))
	fmt.Println("part 2 -", currentDeletePretendent)
}
