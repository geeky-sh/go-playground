package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	tick := time.NewTicker(500 * time.Millisecond)

	go func() {
		for {
			select {
			case <-done: // this is how to check whether the channel is closed
				fmt.Println("End of Goroutine")
				return

			case t := <-tick.C:
				fmt.Printf("Ping at %s\n", t)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	tick.Stop()
	done <- true
	// close(done)
	//fmt.Println("End of Main") // this is used so that go routine can also executed till the time this is printed
}
