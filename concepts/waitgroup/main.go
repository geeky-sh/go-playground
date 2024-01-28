package main

import (
	"fmt"
	"sync"
	"time"
)

func workerMultiply(ch chan int, i int) {
	fmt.Printf("Worker %d started\n", i)
	time.Sleep(1 * time.Second)
	ch <- i * 10
	fmt.Printf("Worker %d finished\n", i)
}

func parentWithoutWg() {
	iter := 100
	ch := make(chan int, iter)

	for i := 1; i <= iter; i++ {
		go workerMultiply(ch, i)
	}

	for i := 1; i <= iter; i++ {
		fmt.Println(<-ch)
	}
}

func parentWithWg() {
	iter := 100
	ch := make(chan int, iter)
	var wg sync.WaitGroup

	for i := 1; i <= iter; i++ {
		wg.Add(1)
		i := i
		go func() {
			workerMultiply(ch, i)
			defer wg.Done()
		}()
	}

	wg.Wait()

	for i := 1; i <= iter; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	parentWithWg()
}
