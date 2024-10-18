package main

import (
	"fmt"
	"github.com/gwkeo/Tower_Back_HW/hw2/queue"
)

func main() {
	q := queue.Queue{}
	fmt.Println(q.IsExist(1))
	fmt.Println("----------------------")
	q.Add(1)
	q.Add(2)
	q.Add(3)
	fmt.Println(q.IsExist(1))
	fmt.Println(q.IsExist(2))
	fmt.Println(q.IsExist(3))
	fmt.Println("----------------------")
	val := q.Pop()
	fmt.Println(q.IsExist(1), val)
	fmt.Println(q.IsExist(2))
	fmt.Println(q.IsExist(3))

}
