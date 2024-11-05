package main

import (
	"fmt"
	"time"
)

func Second() {
	arr := []int{2, 4, 6, 8, 10}
	for _, v := range arr {
		go fmt.Println(v * v)
	}
}

func Third() {
	arr := []int{2, 4, 6, 8, 10}
	sum := 0
	for _, v := range arr {
		sum += v * v
	}
	fmt.Println("sum =", sum)
}

func main() {
	go Second()
	go Third()
	time.Sleep(time.Second)
}
