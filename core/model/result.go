package model

type Result struct {
	Seed      *Seed
	ResultMap map[string]string
}

func NewResult(seed *Seed) *Result {
	result := make(map[string]string)
	return &Result{Seed: seed, ResultMap: result}
}

func (result *Result) GetSeed() *Seed {
	return result.Seed
}

func (result *Result) AddResult(key string, value string) {
	result.ResultMap[key] = value
}

func (result *Result) GetResult(key string) (string, bool) {
	r, ok := result.ResultMap[key]
	return r, ok
}

func (result *Result) GetAllResult() map[string]string {
	return result.ResultMap
}
