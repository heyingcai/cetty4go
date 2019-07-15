package handler

import "github.com/heyingcai/getty/core/crawler"

type Context interface {
	Getty() *crawler.Getty
	HandlerPipeline() *Pipeline
	ProcessHandler() *ProcessHandler
	ReduceHandler() *ReduceHandler
	FireReceive()
	FireDownload()
	FireProcess()
	FireReduce()
}
