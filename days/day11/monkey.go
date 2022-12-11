package main

import (
	"fmt"
)

type monkey struct {
	items        []uint64
	expression   string
	test         uint64
	conditions   map[bool]int // map condition to monkey index
	inspectCount int
}

func newMonkey() *monkey {
	return &monkey{
		items:      make([]uint64, 0, 10),
		conditions: make(map[bool]int, 2),
	}
}

func (m *monkey) addItem(item uint64) *monkey {
	fmt.Println("adding item", item)
	m.items = append(m.items, item)
	return m
}

func (m *monkey) setExpression(expression string) *monkey {
	m.expression = expression

	return m
}

func (m *monkey) setTest(testCondition uint64) *monkey {
	m.test = testCondition
	return m
}
