package api

type Third struct {
}

func NewThird() *Third {
	return &Third{}
}

func (f *Third) GetThird() string {
	return "third"
}
