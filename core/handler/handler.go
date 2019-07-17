package handler

type Handler interface {
	Process()
	Reduce()
}

type ProcessHandler struct {
}

func (handler *ProcessHandler) Process() {

}

type ReduceHandler struct {
}

func (handler *ReduceHandler) Reduce() {

}
