package server

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"

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
	log     *logrus.Logger
}

func NewServer(cpu1, cpu2 *cpu.CPU, buf *buffer.Buffer, endTime time.Time, logger *logrus.Logger) *Server {
	return &Server{
		Cpu1:    cpu1,
		Cpu2:    cpu2,
		Buff:    buf,
		EndTime: endTime,
		log:     logger,
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
		currentQuery, err := s.Buff.GetQuery()
		if err != nil {
			s.log.Info("Buffer is empty, queries not found")
			continue
		}

		if s.Cpu1.IsBroken() && s.Cpu2.IsBroken() {
			s.log.Info("Both CPUs are broken, waiting...")
			continue
		}

		if s.Cpu1.IsBroken() || s.Cpu2.IsBroken() {
			s.log.Info("One of CPU is broken, decreasing handling speed...")
			handleFunc = slowHandle
		} else {
			handleFunc = normalHandle
		}

		s.log.Infof("Server has handled query: ID = [%d], Priority = [%d]", currentQuery.ID, currentQuery.Priority)
		handleFunc(currentQuery)
	}
}

func (s *Server) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	s.log.Infof("Server started: %s", time.Now().String())
	s.HandleQueries()
	s.log.Infof("Server finished: %s", time.Now().String())
}
