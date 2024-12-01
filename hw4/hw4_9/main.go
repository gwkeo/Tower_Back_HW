package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{4, 2, 34, 3, 4, 65, 5, 2, 4, 6, 286, 2626, 2622}
	xc := make(chan int)
	sqrc := make(chan int)

	go func() {
		for _, v := range arr {
			xc <- v
		}
		close(xc)
	}()
	go func() {
		for c := range xc {
			sqrc <- c * c
		}
	}()

	go func() {
		for c := range sqrc {
			fmt.Print(c, "\t")
		}
		close(sqrc)
	}()

	time.Sleep(time.Second)
}
