package main

import (
	"fmt"
	"math/big"
)

type monkey struct {
	items        []*big.Int
	expression   string
	test         *big.Int
	conditions   map[bool]int // map condition to monkey index
	inspectCount int
}

func newMonkey() *monkey {
	return &monkey{
		items:      make([]*big.Int, 0, 10),
		conditions: make(map[bool]int, 2),
	}
}

func (m *monkey) addItem(item *big.Int) *monkey {
	fmt.Println("adding item", item)
	m.items = append(m.items, item)
	return m
}

func (m *monkey) setExpression(expression string) *monkey {
	m.expression = expression

	return m
}

func (m *monkey) setTest(testCondition *big.Int) *monkey {
	m.test = testCondition
	return m
}
