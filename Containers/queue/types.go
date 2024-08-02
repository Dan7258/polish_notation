package queue

type Queue interface {
	Enqueue(data string)
	Dequeue() string
	Size() int
	IsEmpty() bool
}

type Queuelist struct {
	data string
	next *Queuelist
}
