package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseMonkeys(source io.Reader) []*monkey {
	s := bufio.NewScanner(source)
	m := newMonkey()
	ms := []*monkey{
		m,
	}
	idx := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			m = newMonkey()
			ms = append(ms, m)
			idx += 1
			continue
		}

		if strings.HasPrefix(line, "Monkey") {
			continue
		}

		line = strings.Replace(line, ",", "", -1)
		line = strings.Trim(line, " ")
		attributes := strings.Split(line, ":")
		attribute, value := attributes[0], strings.Trim(attributes[1], " ")
		switch attribute {
		case "Starting items":
			for _, itm := range strings.Split(value, " ") {
				item, _ := strconv.ParseUint(itm, 10, 64)
				m.addItem(item)
			}
		case "Operation":
			m.setExpression(strings.Split(value, " = ")[1])
		case "Test":
			rawTestVal := strings.Split(value, " ")
			testVal, _ := strconv.ParseUint(rawTestVal[len(rawTestVal)-1], 10, 64)
			m.setTest(testVal)
		case "If true":
			rawTrueCond := strings.Split(value, " ")
			trueCond, _ := strconv.Atoi(rawTrueCond[len(rawTrueCond)-1])
			m.conditions[true] = trueCond
		case "If false":
			rawFalseCond := strings.Split(value, " ")
			falseCond, _ := strconv.Atoi(rawFalseCond[len(rawFalseCond)-1])
			m.conditions[false] = falseCond
		}
	}

	return ms
}

func getOverlap(ms []*monkey) uint64 {
	result := ms[0].test
	ms = ms[1:]
	for _, m := range ms {
		result *= m.test
	}

	return result
}

func solveP1(source io.Reader, rounds int, worryLevelDiv uint64) int {
	ms := parseMonkeys(source)
	overlap := getOverlap(ms)
	for i := 1; i <= rounds; i++ {
		for _, m := range ms {
			for len(m.items) != 0 {
				item := m.items[0]
				m.items = m.items[1:]
				expr := strings.Split(m.expression, " ")
				operand := expr[1]
				rawArg2 := expr[2]

				var (
					arg1, arg2 uint64
				)

				arg1 = item

				if rawArg2 == "old" {
					arg2 = item
				} else {
					arg2, _ = strconv.ParseUint(rawArg2, 10, 64)
				}

				var newItem uint64
				switch operand {
				case "*":
					newItem = arg1 * arg2
				case "+":
					newItem = arg1 + arg2
				}

				if worryLevelDiv == 1 {
					newItem %= overlap
				} else {
					newItem /= 3
				}

				mIdx := m.conditions[newItem%m.test == 0]
				ms[mIdx].items = append(ms[mIdx].items, newItem)
				m.inspectCount++
			}
		}
	}

	topInspectingMonkeys := [2]int{0, 0}
	for _, m := range ms {
		if m.inspectCount > topInspectingMonkeys[1] && m.inspectCount > topInspectingMonkeys[0] {
			topInspectingMonkeys[0] = topInspectingMonkeys[1]
			topInspectingMonkeys[1] = m.inspectCount
			continue
		}

		if m.inspectCount > topInspectingMonkeys[0] {
			topInspectingMonkeys[0] = m.inspectCount
		}
	}

	return topInspectingMonkeys[0] * topInspectingMonkeys[1]
}

func main() {
	{
		f, err := os.Open("days/day11/input.txt")
		defer f.Close()
		if err != nil {
			log.Fatalln(err)
		}

		solveP1(f, 20, 3)
	}
	{
		f, err := os.Open("days/day11/input.txt")
		defer f.Close()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(solveP1(f, 10000, 1))
	}
}
