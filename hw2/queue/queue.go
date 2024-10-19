package queue

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
	prevNode := q.head
	q.head = q.head.next
	prevNode.next = nil
	return prevNode.val
}

func (q *Queue) IsExist(val int) bool {
	n := q.head
	for ; n != nil; n = n.next {
		if n.val == val {
			return true
		}
	}
	return false
}
