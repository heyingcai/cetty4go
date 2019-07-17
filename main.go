package main

import "github.com/heyingcai/getty/core/handler"

func main() {

	processHandler := new(handler.ProcessHandler)

	reduceHandler := new(handler.ReduceHandler)

	processHandler.Process()

	reduceHandler.Reduce()
}
