package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name]++
}

func main() {

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doInc := func(name string, n int) {
		for range n {
			c.inc(name)
		}

		wg.Done()
	}

	wg.Add(3)

	go doInc("a", 200000)
	go doInc("b", 50000)
	go doInc("a", 6000)

	wg.Wait()
	fmt.Println(c.counters)

}
