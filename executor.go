package cron

import (
	"sync"
)

type defaultExecutor struct {
}

func (executor defaultExecutor) exec(jobWaiter *sync.WaitGroup, job Job) {
	jobWaiter.Add(1)
	go func() {
		defer jobWaiter.Done()
		job.Run()
	}()
}

// If you want to use the goroutine pool, the caller implements it by himself
//type GoroutinePoolExecutor struct {
//	Pool		*ants.Pool
//}
//
//func (executor *GoroutinePoolExecutor) exec(job Job) {
//	executor.jobWaiter.Add(1)
//	fun := func() {
//		defer executor.jobWaiter.Done()
//		job.Run()
//	}
//	if err := executor.Pool.Submit(fun); err != nil {
//		executor.jobWaiter.Done()
//	}
//}
//
//func NewGoroutinePoolExecutor(cron *Cron, poolSize int) JobExecutor {
//	pool, err := ants.NewPool(poolSize)
//	if err != nil {
//		panic(err)
//	}
//	return &GoroutinePoolExecutor{jobWaiter: &cron.jobWaiter, Pool: pool}
//}
