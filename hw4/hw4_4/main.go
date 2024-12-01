package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const workersNum = 3

func worker(ctx context.Context, wg *sync.WaitGroup, id int, jobs <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d stopped\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				log.Printf("worker %d stopped\n", id)
			}
			log.Printf("worker %d doing job %d\n", id, job)
			time.Sleep(1 * time.Second)
			log.Printf("worker %d done job %d\n", id, job)
		}
	}
}

func main() {
	jobs := make(chan int)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		k := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				jobs <- k
				k++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i, jobs)
	}

	wg.Wait()
	log.Println("all workers has done their jobs, shutting down... ")
}
