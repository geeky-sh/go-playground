package main

import (
	"fmt"
	"sync"
)

type Cont struct {
	mx sync.Mutex
	c  int
	nc int
}

func (r *Cont) inc() {
	// use mutex to increment the value
	r.mx.Lock()
	defer r.mx.Unlock()
	r.c++
}

func (r *Cont) ncinc() {
	// increment the value as is
	r.nc++
}

func do() {
	var wg sync.WaitGroup
	c := Cont{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(c *Cont) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				c.inc()
				c.ncinc()
			}
		}(&c)
	}

	wg.Wait()

	fmt.Printf("C: %d\n", c.c)
	fmt.Printf("NC: %d\n", c.nc)
}

func main() {
	do()
}
