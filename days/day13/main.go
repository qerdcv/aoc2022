package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
)

type pair struct {
	val, nesting int
	isLastInNest bool
}

func writePairsToCh(source io.Reader, pairsCh chan<- [2]string) {
	defer close(pairsCh)

	pairs := [2]string{
		"",
		"",
	}

	idx := 0
	s := bufio.NewScanner(source)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}

		pairs[idx] = line
		idx++
		if idx == 2 {
			pairsCh <- pairs
			idx = 0
		}
	}
}

func parsePairToCh(ctx context.Context, p string, out chan<- pair) {
	lastCh := ' '
	num := ""
	nesting := 0
	isLastInNest := false

	defer close(out)

	for _, ch := range p {
		switch ch {
		case '[':
			isLastInNest = false
			nesting++
		case ']', ',':
			if lastCh == ']' {
				continue
			}

			if ch == ']' {
				isLastInNest = true
			}

			i, _ := strconv.Atoi(num)
			select {
			case <-ctx.Done():
				return
			case out <- pair{val: i, nesting: nesting, isLastInNest: isLastInNest}:
				num = ""
			}

			if ch == ']' {
				nesting--
			}
		default:
			num += string(ch)
		}
		lastCh = ch
	}
}

// -1 == left  >  right
// 0 == left  == right
// 1 == left  <  right
func cmpPair(left, right pair) int8 {
	if left.nesting == right.nesting {
		if left.val < right.val {
			return 1
		}

		if left.val == right.val {
			if left.isLastInNest && !right.isLastInNest {
				return 1
			}

			return 0
		}
	} else {
		if left.val < right.val || (left.val == right.val && right.nesting > left.nesting) {
			return 1
		}
	}

	return -1
}

func processPair(pairs [2]string) bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lCh, rCh := make(chan pair), make(chan pair)

	go parsePairToCh(ctx, pairs[0], lCh)
	go parsePairToCh(ctx, pairs[1], rCh)

	for {
		left := <-lCh
		right := <-rCh

		switch cmpPair(left, right) {
		case 1:
			return true
		case 0:
			continue
		}

		break
	}

	return false
}
func processPairs(pairsCh <-chan [2]string) uint64 {
	var pairIdxCnt uint64 = 0
	var pairIdx uint64 = 0
	wg := new(sync.WaitGroup)

	for pairs := range pairsCh {
		wg.Add(1)
		pairIdx += 1
		go func(wg *sync.WaitGroup, p [2]string, pIdx uint64) {
			defer wg.Done()
			if processPair(p) {
				atomic.AddUint64(&pairIdxCnt, pIdx)
			}
		}(wg, pairs, pairIdx)
	}

	wg.Wait()
	return pairIdxCnt
}

func solveP1(source io.Reader) uint64 {
	pairsCh := make(chan [2]string)
	go writePairsToCh(source, pairsCh)
	return processPairs(pairsCh)
}

func solveP2(source io.Reader) int {
	pairsCh := make(chan [2]string)
	go writePairsToCh(source, pairsCh)
	var pairs []string
	for p := range pairsCh {
		pairs = append(pairs, p[0], p[1])
	}

	pairs = append(pairs, "[[2]]", "[[6]]")

	sort.Slice(pairs, func(i, j int) bool {
		return processPair([2]string{
			pairs[i],
			pairs[j],
		})
	})

	result := 1

	for idx, p := range pairs {
		if p == "[[2]]" || p == "[[6]]" {
			result *= idx + 1
		}
	}

	return result
}

func main() {
	{
		f, err := os.Open("days/day13/input.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(solveP1(f))
	}
	{
		f, err := os.Open("days/day13/input.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(solveP2(f))
	}
}
