package main

import (
	"github.com/seaung/crawler/pkg/engine"
	"github.com/seaung/crawler/pkg/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
