package api

type Fifth struct {
}

func NewFifth() *Fifth {
	return &Fifth{}
}

func (f *Fifth) GetFifth() string {
	return "fifth"
}
