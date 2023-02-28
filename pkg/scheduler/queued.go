package scheduler

import (
	"fmt"

	"github.com/seaung/crawler/pkg/engine"
)

type QueuedScheduler struct {
	requestChannel chan engine.Request
	workChannel    chan chan engine.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(request engine.Request) {
	q.requestChannel <- request
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workChannel <- w
}

func (q *QueuedScheduler) Run() {
	q.requestChannel = make(chan chan engine.Request)
	q.workChannel = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeWorker chan engine.Request
			var activeRequest engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-q.requestChannel:
				requestQ = append(requestQ, r)
			case w := <-q.workChannel:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
				fmt.Printf("Request Q : %d, Worker Q : %d\n", len(requestQ), len(workerQ))
			}
		}
	}()
}
