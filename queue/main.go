package main

import (
	"container/list"
	"fmt"
)

// === Implement Queue Using Slices
//func enqueue(queue []int, element int) []int {
//	queue = append(queue, element)
//	fmt.Println("Enqueue:", element)
//	return queue
//}
//
//func dequeue(queue []int) (int, []int) {
//	element := queue[0] // The first element is the one to be dequeued.
//	if len(queue) == 1 {
//		return element, []int{}
//	}
//	return element, queue[1:] // Slice off the element once it is dequeued.
//}
//
//func main() {
//	var queue = make([]int, 0)
//	queue = enqueue(queue, 10)
//	fmt.Println("After pushing 10", queue)
//
//	queue = enqueue(queue, 20)
//	fmt.Println("After pushing 20", queue)
//
//	queue = enqueue(queue, 30)
//	fmt.Println("After pushing 30", queue)
//
//	ele, queue := dequeue(queue)
//	fmt.Println("After removing", ele, queue)
//
//	queue = enqueue(queue, 40)
//	fmt.Println("After pushing 40", queue)
//}

// === Implement Queue Using Structures
//type Queue struct {
//	Elements []int
//	Size     int
//}
//
//func (q *Queue) GetLength() int {
//	return len(q.Elements)
//}
//
//func (q *Queue) IsEmpty() bool {
//	return q.GetLength() == 0
//}
//
//func (q *Queue) Peek() (int, error) {
//	if q.IsEmpty() {
//		return 0, errors.New("empty queue")
//	}
//	return q.Elements[0], nil
//}
//
//func (q *Queue) Enqueue(elem int) {
//	if q.GetLength() == q.Size {
//		fmt.Println("Overflow")
//		return
//	}
//	q.Elements = append(q.Elements, elem)
//}
//
//func (q *Queue) Dequeue() int {
//	if q.IsEmpty() {
//		fmt.Println("Underflow")
//		return 0
//	}
//
//	element := q.Elements[0]
//
//	if q.GetLength() == 1 {
//		q.Elements = nil
//		return element
//	}
//
//	q.Elements = q.Elements[1:]
//	return element
//}
//
//func main() {
//	queue := Queue{Size: 2}
//	fmt.Println(queue.Elements)
//
//	queue.Enqueue(1)
//	fmt.Println(queue.Elements)
//
//	queue.Enqueue(2)
//	fmt.Println(queue.Elements)
//
//	queue.Enqueue(3)
//	fmt.Println(queue.Elements)
//
//	queue.Dequeue()
//	fmt.Println(queue.Elements)
//
//	queue.Dequeue()
//	fmt.Println(queue.Elements)
//
//	queue.Dequeue()
//	fmt.Println(queue.Elements)
//
//	queue.Enqueue(4)
//	fmt.Println(queue.Elements)
//}

// Implement Queue Using LinkList
func main() {
	queue := list.New()

	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	front := queue.Front()
	fmt.Println(front.Value)
	queue.Remove(front)
}
