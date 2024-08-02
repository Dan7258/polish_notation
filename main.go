package main

import (
	"fmt"
	"polish_notation/Containers/queue"
	"polish_notation/Containers/stack"
)

func main() {
	fmt.Printf("goooo\n");
	q := new(queue.Queuelist)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	fmt.Printf("%s\n", q.Dequeue())
	fmt.Println(q.IsEmpty())
	fmt.Printf("%s\n", q.Dequeue())
	fmt.Printf("%s\n", q.Dequeue())
	fmt.Printf("%s\n", q.Dequeue())
	fmt.Println(q.IsEmpty())

	s := new(stack.Stacklist)
	s.Push("a")
	s.Push("b")
	s.Push("c")
	fmt.Printf("%s\n", s.Pop())
	fmt.Printf("%s\n", s.Pop())
	fmt.Printf("%s\n", s.Pop())
	fmt.Printf("%s\n", s.Pop())
}