package main

import (
	"fmt"

	"github.com/chonlatee/simpleds/queue"
	"github.com/chonlatee/simpleds/stack"
)

func main() {
	mystack := stack.NewStack(10)
	fmt.Println(mystack.Size())
	mystack.Push(10)
	fmt.Println(mystack.Size())
	fmt.Println(mystack.Peek())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Size())
	fmt.Println(mystack.Pop())

	myqueue := queue.NewQueue(10)
	myqueue.EnQueue(4)
	myqueue.EnQueue(2)
	myqueue.EnQueue(7)

	fmt.Println(myqueue.DeQueue())
	fmt.Println(myqueue.DeQueue())
	fmt.Println(myqueue.Peek())
	fmt.Println(myqueue.IsEmpty())
	fmt.Println(myqueue.IsFull())
}
