package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

const (
	workersNum      = 3
	jobsArrangement = 100
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int)
	results := make(chan int)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context canceled")
				return
			default:
				jobs <- rand.Intn(jobsArrangement)
			}
		}
	}()

	go func() {
		for i := 0; i < workersNum; i++ {
			select {
			case <-c:
				cancel()
			default:
				go worker(i, jobs, results)
			}
		}
	}()
	time.Sleep(time.Second)
}
