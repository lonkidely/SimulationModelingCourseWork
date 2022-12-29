package main

import (
	"SimulationModelingCourseWork/pkg/logger"
	"github.com/sirupsen/logrus"
	"sync"
	"time"

	"SimulationModelingCourseWork/internal"
	"SimulationModelingCourseWork/internal/buffer"
	"SimulationModelingCourseWork/internal/cpu"
	"SimulationModelingCourseWork/internal/query"
	"SimulationModelingCourseWork/internal/server"
	"SimulationModelingCourseWork/internal/utils"
)

const (
	CpuOneLogPath = "./logs/cpu1"
	CpuTwoLogPath = "./logs/cpu2"
	ServerLogPath = "./logs/server"
	BufferLogPath = "./logs/buffer"
)

func main() {
	endTime := time.Now().Add(2 * time.Second)

	cpu1Logger, closeResp := logger.NewLogger(CpuOneLogPath)
	defer func(closer func() error, log *logrus.Logger) {
		err := closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResp, cpu1Logger)

	cpu2Logger, closeResp := logger.NewLogger(CpuTwoLogPath)
	defer func(closer func() error, log *logrus.Logger) {
		err := closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResp, cpu2Logger)

	serverLogger, closeResp := logger.NewLogger(ServerLogPath)
	defer func(closer func() error, log *logrus.Logger) {
		err := closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResp, serverLogger)

	bufferLogger, closeResp := logger.NewLogger(BufferLogPath)
	defer func(closer func() error, log *logrus.Logger) {
		err := closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResp, bufferLogger)

	cpu1 := cpu.NewCpu(cpu1Logger)
	cpu2 := cpu.NewCpu(cpu2Logger)
	buff := buffer.NewBuffer(bufferLogger)
	serv := server.NewServer(cpu1, cpu2, buff, endTime, serverLogger)

	breakCpu := func(cpu *cpu.CPU, wg *sync.WaitGroup) {
		defer wg.Done()
		for time.Now().Unix() < endTime.Unix() {
			time.Sleep(time.Duration(utils.Exponential(internal.T8)))
			cpu.Break()
			time.Sleep(time.Duration(utils.Normal(internal.T9, internal.T10)))
			cpu.Repair()
		}
	}

	genQuery1 := func(server *server.Server, wg *sync.WaitGroup) {
		defer wg.Done()

		currentID := 1
		for time.Now().Unix() < endTime.Unix() {
			time.Sleep(time.Duration(utils.Uniform(internal.T1, internal.T2)))
			currentQuery := query.NewQuery(currentID, 1)
			currentID++
			server.AddQuery(*currentQuery)
		}
	}

	genQuery2 := func(server *server.Server, wg *sync.WaitGroup) {
		defer wg.Done()

		currentID := 1
		for time.Now().Unix() < endTime.Unix() {
			time.Sleep(time.Duration(utils.Exponential(internal.T3)))
			currentQuery := query.NewQuery(currentID, 2)
			currentID++
			server.AddQuery(*currentQuery)
		}
	}

	genQuery3 := func(server *server.Server, wg *sync.WaitGroup) {
		defer wg.Done()

		currentID := 1
		for time.Now().Unix() < endTime.Unix() {
			time.Sleep(time.Duration(utils.Exponential(internal.T4)))
			currentQuery := query.NewQuery(currentID, 3)
			currentID++
			server.AddQuery(*currentQuery)
		}
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go serv.Start(wg)

	wg.Add(1)
	go breakCpu(cpu1, wg)

	wg.Add(1)
	go breakCpu(cpu2, wg)

	wg.Add(1)
	go genQuery1(serv, wg)

	wg.Add(1)
	go genQuery2(serv, wg)

	wg.Add(1)
	go genQuery3(serv, wg)

	wg.Wait()
}
