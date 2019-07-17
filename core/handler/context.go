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

type GenericHandlerContext struct {
	ProcessEvent bool
	ReduceEvent  bool
	Name         string
	Prev         *GenericHandlerContext
	Next         *GenericHandlerContext
	Pipeline     *Pipeline
}

func NewAbstractHandlerContext(processEvent bool, reduceEvent bool, name string, prev *GenericHandlerContext, next *GenericHandlerContext, handlerPipeline *Pipeline) *GenericHandlerContext {
	return &GenericHandlerContext{ProcessEvent: processEvent, ReduceEvent: reduceEvent, Name: name, Prev: prev, Next: next, Pipeline: handlerPipeline}
}

func (handler *GenericHandlerContext) Getty() *crawler.Getty {
	return handler.Pipeline.Getty
}

func (handler *GenericHandlerContext) HandlerPipeline() *Pipeline {
	return handler.Pipeline
}

func (handler *GenericHandlerContext) Named() string {
	return handler.Name
}
