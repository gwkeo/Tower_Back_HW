package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func generator(jobs chan int) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	defer os.Exit(1)
	for i := 0; ; i++ {
		select {
		case <-signals:
			fmt.Println("Got SIGINT/SIGTERM")
			close(jobs)
			return
		default:
			jobs <- i
		}
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	go generator(jobs)
	for a := 1; ; a++ {
		<-results
	}
}
