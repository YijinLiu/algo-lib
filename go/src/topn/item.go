package topn

type Item interface {
	Key() string
}

type StringItem string

func (s StringItem) Key() string {
	return string(s)
}
