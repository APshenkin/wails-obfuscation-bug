package api

type Fourth struct {
}

func NewFourth() *Fourth {
	return &Fourth{}
}

func (f *Fourth) GetFourth() string {
	return "fourth"
}
