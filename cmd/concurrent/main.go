package main

import (
	"github.com/seaung/crawler/pkg/engine"
	"github.com/seaung/crawler/pkg/scheduler"
	"github.com/seaung/crawler/pkg/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 50,
	}

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
