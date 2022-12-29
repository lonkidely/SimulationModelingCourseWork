package cpu

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type CPU struct {
	mu      *sync.Mutex
	Healthy bool
	log     *logrus.Logger
}

func NewCpu(logger *logrus.Logger) *CPU {
	return &CPU{
		mu:      &sync.Mutex{},
		Healthy: true,
		log:     logger,
	}
}

func (c *CPU) Break() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.log.Info("I have been broken, sorry :(")
	c.Healthy = false
}

func (c *CPU) Repair() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.log.Info("Make CPU great again!")
	c.Healthy = true
}

func (c *CPU) IsBroken() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.Healthy
}
