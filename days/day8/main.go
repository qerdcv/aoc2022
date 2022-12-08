package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func parseTreeMap(source io.Reader) [][]int {
	s := bufio.NewScanner(source)
	var treeMap [][]int
	for s.Scan() {
		line := s.Text()
		treeLine := make([]int, 0, len(line))
		for _, tree := range line {
			treeLine = append(treeLine, int(tree-48)) // ascii to int
		}
		treeMap = append(treeMap, treeLine)
	}

	return treeMap
}

func isVisible(tree int, treeMap [][]int, row, col int) bool {
	// vertical checks
	// check from top
	topCheck := false
	botCheck := false
	rightCheck := false
	leftCheck := false

	for i := 0; i < row; i++ {
		if treeMap[i][col] >= tree {
			topCheck = true
			break
		}
	}

	// check to bot
	for i := len(treeMap) - 1; i > row; i-- {
		if treeMap[i][col] >= tree {
			botCheck = true
			break
		}
	}

	// check horizontal
	// check to left
	for i := 0; i < col; i++ {
		if treeMap[row][i] >= tree {
			leftCheck = true
			break
		}
	}

	// check to right
	for i := len(treeMap[row]) - 1; i > col; i-- {
		if treeMap[row][i] >= tree {
			rightCheck = true
			break
		}
	}

	return topCheck && botCheck && rightCheck && leftCheck
}

func calcTreeScore(tree int, treeMap [][]int, row, col int) int {
	// vertical checks
	// check from top
	topUnblockedCnt := 0
	botUnblockedCnt := 0
	leftUnblockedCnt := 0
	rightUnblockedCnt := 0

	for i := row - 1; i >= 0; i-- {
		topUnblockedCnt++
		if treeMap[i][col] >= tree {
			break
		}
	}

	// check to bot
	for i := row + 1; i < len(treeMap); i++ {
		botUnblockedCnt++
		if treeMap[i][col] >= tree {
			break
		}
	}

	// check horizontal
	// check to left
	for i := col - 1; i >= 0; i-- {
		leftUnblockedCnt++
		if treeMap[row][i] >= tree {
			break
		}
	}

	// check to right
	for i := col + 1; i < len(treeMap[row]); i++ {
		rightUnblockedCnt++
		if treeMap[row][i] >= tree {
			break
		}
	}

	return topUnblockedCnt * botUnblockedCnt * rightUnblockedCnt * leftUnblockedCnt
}

func calcVisibleTrees(treeMap [][]int) int {
	visibleTrees := len(treeMap)*2 + (len(treeMap[0])*2 - 4)
	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap[i])-1; j++ { // iterate over inner trees
			currentTree := treeMap[i][j]
			if !isVisible(currentTree, treeMap, i, j) {
				visibleTrees += 1
			}
		}
	}

	return visibleTrees
}

func calcTreesScore(treeMap [][]int) int {
	maxScore := 0
	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap[i])-1; j++ { // iterate over inner trees
			if newScore := calcTreeScore(treeMap[i][j], treeMap, i, j); newScore > maxScore {
				maxScore = newScore
			}
		}
	}

	return maxScore
}

func main() {
	f, err := os.Open("days/day8/input.txt")
	if err != nil {
		panic(err)
	}

	treeMap := parseTreeMap(f)
	fmt.Println("visible trees: ", calcVisibleTrees(treeMap))
	fmt.Println("max tree score:", calcTreesScore(treeMap))
}
