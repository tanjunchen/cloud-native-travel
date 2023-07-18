package v3

import (
	"go-learn/spider/zhenaiwang/v3/engine"
	"go-learn/spider/zhenaiwang/v3/parser"
	"go-learn/spider/zhenaiwang/v3/scheduler"
)

func Start() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
