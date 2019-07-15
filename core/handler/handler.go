package handler

type Handler interface {
	Reduce()
}

type ProcessHandler struct {
}

type ReduceHandler struct {
}
