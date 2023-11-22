package api

type Second struct {
}

func NewSecond() *Second {
	return &Second{}
}

func (f *Second) GetSecond() string {
	return "second"
}
