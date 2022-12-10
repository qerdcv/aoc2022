package main

type computer struct {
	*crt
	xRegister, signalStrength, clockCycles int
}

func newComputer() computer {
	return computer{
		crt:       new(crt),
		xRegister: 1,
	}
}

func (c *computer) noop() {
	c.clockTick()
}

func (c *computer) addX(val int) {
	c.clockTick()
	c.clockTick()
	c.xRegister += val
}

func (c *computer) clockTick() {
	c.draw(c.xRegister)
	c.clockCycles += 1
	if (c.clockCycles-20)%40 == 0 {
		c.signalStrength += c.clockCycles * c.xRegister
	}
}
