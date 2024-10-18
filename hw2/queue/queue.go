package queue

import "fmt"

type Node struct {
	next *Node
	val  int
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) Add(val int) {
	newNode := &Node{val: val}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue) Pop() int {
	val := q.head.val
	q.head = q.head.next
	q.head.next = nil
	return val
}

func (q *Queue) IsExist(val int) bool {
	n := q.head
	for ; n != nil; n = n.next {
		if n.val == val {
			fmt.Println(n.val)
			return true
		}
	}
	return false
}
