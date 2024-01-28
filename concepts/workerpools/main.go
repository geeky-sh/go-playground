package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Worker Pools:
Simular worker pools

What is done here?
1. Create M workers via Goroutines
2. Each worker can take one or more jobs, and put the results in the results channel
3. Print all the results at the end of the function
4. We simulate job run time by waiting for 1 second for job to finish

Observations:
1. Even though channel is shared between workers, go ensures that an item is only provided to one worker at the time.
2. A channel can be shared between multiple goroutines
3. Waitgroup is used to ensure that all the jobs are finished before we print the results and exit main
4. Here, we did state management through worker pools

Imp. piece of code in Function Parameters:
`<- chan int` // receiver
`chan<- int` // sender

ref: https://gobyexample.com/worker-pools
*/

var wg sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("iterate job %d by worker %d\n", job, id)
		result := job * 10
		time.Sleep(1 * time.Second)
		results <- result
		wg.Done()
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	numWorkers := 5
	numJobs := 20

	for i := 1; i <= numWorkers; i++ {
		i := i
		go worker(i, jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		fmt.Printf("create job %d\n", i)
		wg.Add(1)
		jobs <- i
	}
	close(jobs)

	wg.Wait()

	for i := 1; i <= numJobs; i++ {
		fmt.Printf("\nResult for Job %d - %d\n", i, <-results)
	}

}
