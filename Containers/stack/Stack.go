package stack

func (q *Stacklist) Push(data string) {
	if q.data == "" && q.next == nil {
		q.data = data
	} else {
		current := q
		for current.next != nil {
			current = current.next
		}
		newElem := new(Stacklist)
		newElem.data = data
		current.next = newElem
	}
}

func (q *Stacklist) Pop() string {

	if q.data == "" && q.next == nil {
		return ""
	}
	current := q
	precurrent := current
	for current.next != nil {
		precurrent = current
		current = current.next
	}
	data := current.data
	current.data = ""
	precurrent.next = nil

	return data
}

func (q *Stacklist) Clear() {
	q.data = ""
	q.next = nil
}
