package taskrunner

import (
	"time"
	"log"
)

type Worker struct {
	ticket *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		//ticket: time.NewTicker(interval*time.Second),
		ticket: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <- w.ticket.C:
			log.Printf("ticket run start--------------------\n")
			go w.runner.StartAll()
			log.Printf("ticket run end--------------------\n")
		}
	}
}


func Start()  {
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()
}