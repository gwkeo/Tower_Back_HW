package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	runtimeSeconds     = 3
	newValMilliseconds = 200
	randN              = 10
)

func main() {
	c := make(chan int)
	go func() {
		for {
			c <- rand.Intn(randN)
			time.Sleep(newValMilliseconds * time.Millisecond)
		}
	}()

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
	time.Sleep(runtimeSeconds * time.Second)
	close(c)
}
