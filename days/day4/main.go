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

type Union struct {
	leftLim  int
	rightLim int
}

func NewUnionFromRange(u string) Union {
	lims := strings.Split(u, "-")
	rawLeftLim, rawRightLim := lims[0], lims[1]
	leftLim, _ := strconv.Atoi(rawLeftLim)
	rightLim, _ := strconv.Atoi(rawRightLim)
	return Union{
		leftLim:  leftLim,
		rightLim: rightLim,
	}
}

func checkUnion(u1, u2 Union) bool {
	return ((u1.leftLim >= u2.leftLim && u1.rightLim <= u2.rightLim) ||
		(u2.leftLim >= u1.leftLim && u2.rightLim <= u1.rightLim))
}

func calcUnions(source io.ReadCloser) (int, error) {
	unionCnt := 0
	r := bufio.NewReader(source)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, fmt.Errorf("read line: %w", err)
		}
		unions := strings.Split(string(b), ",")
		u1, u2 := NewUnionFromRange(unions[0]), NewUnionFromRange(unions[1])
		if checkUnion(u1, u2) {
			unionCnt++
		}
	}

	return unionCnt, nil
}

func checkUnionOverlap(u1, u2 Union) bool {
	return (u1.leftLim >= u2.leftLim && u1.leftLim <= u2.rightLim) || (u1.rightLim >= u2.leftLim && u1.leftLim <= u2.rightLim)
}

func calcUnionsP2(source io.ReadCloser) (int, error) {
	defer source.Close()

	unionCnt := 0
	r := bufio.NewReader(source)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return 0, fmt.Errorf("read line: %w", err)
		}

		unions := strings.Split(string(b), ",")
		u1, u2 := NewUnionFromRange(unions[0]), NewUnionFromRange(unions[1])
		if checkUnionOverlap(u1, u2) {
			unionCnt++
		}
	}

	return unionCnt, nil
}

func main() {
	f, err := os.Open("days/day4/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(calcUnions(f))
	f.Seek(0, 0)
	fmt.Println(calcUnionsP2(f))
}
