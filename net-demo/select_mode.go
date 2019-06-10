package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	MaxWorker = 1000
	MaxQueue  = 200000
	wg        sync.WaitGroup
)

type PayLoad struct {
}

func (p PayLoad) do() {
	time.Sleep(time.Millisecond * 200)
	wg.Done()
}

type Job struct {
	payLoad PayLoad
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				job.payLoad.do()
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	// 开始运行 n 个 worker
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}

func main() {
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	JobQueue = make(chan Job, MaxQueue)
	fmt.Println("start process")
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		payload := PayLoad{}
		work := Job{payLoad: payload}
		JobQueue <- work
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
