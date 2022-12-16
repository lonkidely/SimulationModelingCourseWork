package cpu

import "sync"

type CPU struct {
	mu      *sync.Mutex
	Healthy bool
}

func NewCpu() *CPU {
	return &CPU{
		mu:      &sync.Mutex{},
		Healthy: true,
	}
}

func (c *CPU) Break() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Healthy = false
}

func (c *CPU) Repair() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Healthy = true
}

func (c *CPU) IsBroken() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.Healthy
}
