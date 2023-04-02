package engine

import "log"

// 并发引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// 任务调度器
type Scheduler interface {
	Submit(request Request) // 提交任务
	ConfigMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigMasterWorkerChan(in)

	// 创建 goroutine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	itemCount := 0
	for {
		// 接受 Worker 的解析结果
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: #%d: %v\n", itemCount, item)
			itemCount++
		}

		// 然后把 Worker 解析出的 Request 送给 Scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
