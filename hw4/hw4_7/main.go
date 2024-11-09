package main

import (
	"fmt"
	"sync"
)

const (
	iterN = 10000
)

type smt struct {
	mu sync.Mutex
	m  map[string]int
}

func (a *smt) Add(s string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.m[s]++
}

func main() {
	a := smt{
		m: map[string]int{"a": 1, "b": 2, "c": 3},
	}
	wg := sync.WaitGroup{}
	increment := func(name string, n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			a.Add(name)
		}
	}

	wg.Add(3)
	go increment("a", iterN)
	go increment("b", iterN)
	go increment("c", iterN)
	wg.Wait()
	fmt.Println(a.m)
}
