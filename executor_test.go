package cron

import (
	"sync"
	"testing"
	"time"
)

func TestExecutorExec(t *testing.T) {
	executor := defaultExecutor{}
	jobWaiter := sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		executor.exec(&jobWaiter, FuncJob(func() {
			time.Sleep(time.Second)
		}))
	}
	jobWaiter.Wait()
}
