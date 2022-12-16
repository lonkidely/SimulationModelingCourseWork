package cpu

import "sync"

type CPU struct {
	mu       *sync.Mutex
	IsBroken bool
}

func NewCpu() *CPU {
	return &CPU{
		mu:       &sync.Mutex{},
		IsBroken: false,
	}
}

func (c *CPU) Break() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.IsBroken = true
}

func (c *CPU) Repair() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.IsBroken = false
}
