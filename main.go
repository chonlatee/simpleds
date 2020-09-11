package main

import (
	"fmt"

	"github.com/chonlatee/simpleds/linklist"

	"github.com/chonlatee/simpleds/queue"
	"github.com/chonlatee/simpleds/stack"
)

func main() {
	mystack := stack.NewStack(10)
	// fmt.Println(mystack.Size())
	mystack.Push(10)
	// fmt.Println(mystack.Size())
	// fmt.Println(mystack.Peek())
	// fmt.Println(mystack.Pop())
	// fmt.Println(mystack.Size())

	myqueue := queue.NewQueue(10)
	myqueue.EnQueue(4)
	// myqueue.EnQueue(2)
	// myqueue.EnQueue(7)

	// fmt.Println(myqueue.DeQueue())
	// fmt.Println(myqueue.DeQueue())
	// fmt.Println(myqueue.Peek())
	// fmt.Println(myqueue.IsEmpty())
	// fmt.Println(myqueue.IsFull())

	fmt.Println("Linklist")

	myLinklist := linklist.NewSingleLinklist()

	// myLinklist.InsertFront(10)
	// myLinklist.InsertFront(9)
	// myLinklist.InsertFront(8)
	// myLinklist.InsertFront(7)

	// myLinklist.InsertBack(10)
	// myLinklist.InsertBack(9)
	// myLinklist.InsertBack(7)
	// myLinklist.Traverse()
	// myLinklist.InsertBefore(8, 7)
	// myLinklist.Traverse()
	// myLinklist.InsertAfter(100, 10)
	// myLinklist.Traverse()
	// myLinklist.InsertAfter(6, 7)
	// myLinklist.Traverse()
	// myLinklist.RemoveBack()
	// myLinklist.Traverse()
	// fmt.Println()
	// myLinklist.RemoveBefore(7)
	// fmt.Println()
	// myLinklist.Traverse()
	// myLinklist.RemoveAfter(9)
	// myLinklist.Traverse()

	// myLinklist.ShowHeadValue()
	// myLinklist.ShowTailvalue()
	// myLinklist.InsertAfter(8, 9)
	// myLinklist.Traverse()
	// myLinklist.ShowHeadValue()
	// myLinklist.ShowTailvalue()

}
