package server

import (
	"fmt"
	"sync"
	"time"

	"SimulationModelingCourseWork/internal"
	"SimulationModelingCourseWork/internal/buffer"
	"SimulationModelingCourseWork/internal/cpu"
	"SimulationModelingCourseWork/internal/query"
	"SimulationModelingCourseWork/internal/utils"
)

type Server struct {
	Cpu1    *cpu.CPU
	Cpu2    *cpu.CPU
	Buff    *buffer.Buffer
	EndTime time.Time
}

func NewServer(cpu1, cpu2 *cpu.CPU, buf *buffer.Buffer, endTime time.Time) *Server {
	return &Server{
		Cpu1:    cpu1,
		Cpu2:    cpu2,
		Buff:    buf,
		EndTime: endTime,
	}
}

func (s *Server) AddQuery(q query.Query) {
	s.Buff.AddQuery(q)
}

func (s *Server) HandleQueries() {
	slowHandle := func(q query.Query) {
		switch q.Priority {
		case 1:
			time.Sleep(time.Duration(utils.Exponential(internal.T5)) / 2)
		case 2:
			time.Sleep(time.Duration(utils.Exponential(internal.T6)) / 2)
		case 3:
			time.Sleep(time.Duration(utils.Exponential(internal.T7)) / 2)
		}
	}

	normalHandle := func(q query.Query) {
		switch q.Priority {
		case 1:
			time.Sleep(time.Duration(utils.Exponential(internal.T5)))
		case 2:
			time.Sleep(time.Duration(utils.Exponential(internal.T6)))
		case 3:
			time.Sleep(time.Duration(utils.Exponential(internal.T7)))
		}
	}

	var handleFunc func(q query.Query)

	for time.Now().Unix() < s.EndTime.Unix() {
		currentEvent, err := s.Buff.GetQuery()
		if err != nil {
			continue
		}

		if s.Cpu1.IsBroken && s.Cpu2.IsBroken {
			fmt.Println("Both CPUs are broken, waiting...")
			continue
		}

		if s.Cpu1.IsBroken || s.Cpu2.IsBroken {
			fmt.Println("One of CPU is broken, decreasing handling speed...")
			handleFunc = slowHandle
		} else {
			handleFunc = normalHandle
		}

		fmt.Printf("Server has handled query: ID = [%d], Priority = [%d]\n", currentEvent.ID, currentEvent.Priority)
		handleFunc(currentEvent)
	}
}

func (s *Server) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Server started: %s", time.Now().String())
	s.HandleQueries()
	fmt.Printf("Server finished: %s", time.Now().String())
}
