package main

import (
	"strconv"
)

type direction string

const (
	up    direction = "U"
	left  direction = "L"
	down  direction = "D"
	right direction = "R"
)

type RopeNode struct {
	x, y int
	next *RopeNode
}

func (rn *RopeNode) moveTail() {
	head := rn
	tail := rn.next

	dy := head.y - tail.y
	dx := head.x - tail.x

	if abs(dy) > 1 {
		tail.y += dy - sign(dy)
		if abs(dx) == 1 {
			tail.x += dx
		}
	}

	if abs(dx) > 1 {
		tail.x += dx - sign(dx)
		if abs(dy) == 1 {
			tail.y += dy
		}
	}
}

type rope struct {
	head             *RopeNode
	tail             *RopeNode
	visitedPositions map[string]bool
}

func newRope(tailLen int) *rope {
	r := &rope{
		head:             new(RopeNode),
		visitedPositions: map[string]bool{},
	}

	node := r.head
	for i := 0; i < tailLen; i++ {
		node.next = new(RopeNode)
		node = node.next
		r.tail = node
	}

	return r
}

func (r *rope) move(dir direction, delta int) {
	for i := 1; i <= delta; i++ {
		switch dir {
		case up:
			r.head.y--
		case left:
			r.head.x--
		case down:
			r.head.y++
		case right:
			r.head.x++
		}

		for node := r.head; node.next != nil; node = node.next {
			node.moveTail()
		}

		r.visitedPositions[hashPos(r.tail.x, r.tail.y)] = true
	}
}

func (r *rope) getVisitedPositions() int {
	return len(r.visitedPositions)
}

func hashPos(x, y int) string { // dummy hash-function to store uniq position of tail
	return strconv.Itoa(x) + strconv.Itoa(y)
}
