package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Atomic Counters:
Go lang function provided atomic way of updating the value

What is done here:
- We simulate considerable workers updatting the same set of vars
- We used atomic funcs provided by go to update the value for one var and used increment operator for another

Observations:
- We see the one which uses atomic operator gives the corrrect results
- While the one which uses increment operator gives incorrect ones. This is because, since multiple goroutines access the same value,
we get inconsistent results since the same operation is done by multile go routines.

Conclusion:
For state management in the case of goroutines, the ways are:
- Channels
- Atomic Counters
*/

func main() {
	var wg sync.WaitGroup
	var op1 uint64
	var op2 uint64

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				atomic.AddUint64(&op1, 1)
				op2++
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Atomic op1 %d\n", op1)
	fmt.Printf("Non-atomic op2 %d\n", op2)

}
