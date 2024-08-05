package queue

func (q *Queuelist) Enqueue(data string) {
	if q.data == "" && q.next == nil {
		q.data = data
	} else {
		current := q
		for current.next != nil {
			current = current.next
		}
		newElem := new(Queuelist)
		newElem.data = data
		current.next = newElem
	}
}

func (q *Queuelist) Dequeue() string {
	data := q.data
	if q.data == "" && q.next == nil {
		return ""
	}
	if q.next == nil {
		q.data = ""
	} else {
		*q = *q.next
	}
	return data
}

func (q *Queuelist) Size() int {
	current := q
	i := 0
	for ; current.next != nil; current = current.next {
		i++
	}
	return i
}

func (q *Queuelist) IsEmpty() bool {
	f := false
	if q.data == "" && q.next == nil {
		f = true
	}
	return f
}
