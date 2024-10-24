package queue

type Node[T comparable] struct {
	next *Node[T]
	val  T
}

type Queue[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

func (q *Queue[T]) Add(val T) {
	newNode := &Node[T]{val: val}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue[T]) Pop() T {
	currHead := q.head
	q.head = q.head.next
	currHead.next = nil
	return currHead.val
}

func (q *Queue[T]) IsExist(val T) bool {
	if q.head != nil {
		for curr := q.head; curr != nil; curr = curr.next {
			if curr.val == val {
				return true
			}
		}
	}
	return false
}
