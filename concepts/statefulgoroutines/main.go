package main

import (
	"fmt"
	"sync"
)

/*
Maintaining State by Having a separate GoRoutine which updates the value while communicating through Channel.

Obervations:
- Closing the channel refers to everything in the buffered channel emptied
- We can have a separate Goroutine responsible for updating values
- Buffered channels are only needed when events are sent before the receiver is added, else no need.


*/

type Counter struct {
	sf       int
	ntsf     int
	writeReq chan int
	quit     chan int
}

func NewCounter() *Counter {
	c := &Counter{sf: 0, ntsf: 0}
	c.writeReq = make(chan int)
	c.quit = make(chan int)

	go func() {
		for {
			select {
			case <-c.writeReq:
				c.sf++
			case <-c.quit:
				fmt.Println("QUIT")
				return
			}
		}

	}()
	return c
}

func (r *Counter) Close() {
	close(r.writeReq)
	r.quit <- 0
}

func (r *Counter) IncSF(id int) {
	r.writeReq <- id
}

func (r *Counter) IncNSF(id int) {
	r.ntsf++
}

func do() {
	var wg sync.WaitGroup
	c := NewCounter()
	defer c.Close()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			for i := 0; i < 1; i++ {
				c.IncSF(j)
				c.IncNSF(j)
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("SF: %d\n", c.sf)
	fmt.Printf("NSF: %d\n", c.ntsf)
}

func main() {
	do()
}
