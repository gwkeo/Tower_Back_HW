package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	defaultSleepMillis = 800
	mainSleepSeconds   = 3
)

func StopWithBoolChan(c chan bool) {
	for {
		select {
		case <-c:
			fmt.Println("Just stopped with bool chan")
			return
		default:
			fmt.Println("Still running")
			time.Sleep(defaultSleepMillis * time.Millisecond)
		}
	}
}

func StopWithWaitGroup(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Started waitgroup")
	time.Sleep(time.Second * mainSleepSeconds)
}

func StopWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done")
			return
		default:
			fmt.Println("Still running")
			time.Sleep(defaultSleepMillis * time.Millisecond)
		}
	}
}

func sep() { fmt.Println("===============================") }

func main() {
	stopChan := make(chan bool)
	go StopWithBoolChan(stopChan)
	time.Sleep(mainSleepSeconds * time.Second)
	stopChan <- true
	time.Sleep(time.Second)
	close(stopChan)
	fmt.Println("Stop with chanel done")
	sep()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go StopWithWaitGroup(wg)
	wg.Wait()
	fmt.Println("WaitGroup done")
	sep()

	ctx, cancel := context.WithCancel(context.Background())
	go StopWithContext(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
	sep()
}
