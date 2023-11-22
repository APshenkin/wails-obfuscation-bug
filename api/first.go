package api

type First struct {
}

func NewFirst() *First {
	return &First{}
}

func (f *First) GetFirst() string {
	return "first"
}
