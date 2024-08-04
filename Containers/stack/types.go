package stack

type Stack interface {
	Push(data string)
	Pop() string
	Clear()
}

type Stacklist struct {
	data string
	next *Stacklist
}
