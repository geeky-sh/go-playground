package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

/*
I want to verify whether the child goroutine exits when the parent exits when the parent is not the main routine.

Running the code below, I observe that the child process continues their work even after the parent process is completed.
Therefore it is important to ensure that the child process also exists when the parent exits
*/

var wg sync.WaitGroup

func parent() {
	fmt.Println("Parent has started")
	defer fmt.Println("Parent has exited")
	wg.Add(1)
	go child()
	time.Sleep(2 * time.Second)
	wg.Done()
}

func child() {
	fmt.Println("Child has started")
	defer fmt.Println("Child has exited")
	time.Sleep(5 * time.Second)
	wg.Done()
}

func sincereChild(ctx context.Context) {
	fmt.Println("Child has started")
	defer func() {
		fmt.Println("Child has exited")
		wg.Done()
	}()

	f, err := os.Open("test.txt")
	defer func() {
		f.Close()
		fmt.Println("Closing the file")
	}()
	if err != nil {
		log.Fatalf("err in opening file %s\n", err.Error())
	}

	done := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		b := make([]byte, 100)
		n, err := f.Read(b)
		if err != nil {
			log.Fatalf("err in reading file %v\n", err)
		}
		fmt.Println(string(b[:n]))
		fmt.Println("Child has done its work")
		close(done)
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("Child has exited due to context error %v\n", ctx.Err())
		return
	case <-done:
		fmt.Println("Child has done its work 1")
		return
	}
}

func sincereParent() {
	fmt.Println("Parent has started")
	ctx, cncl := context.WithCancel(context.Background())
	defer func() {
		fmt.Println("Parent has exited")
		cncl()
		wg.Done()
	}()
	wg.Add(1)
	go sincereChild(ctx)
	time.Sleep(2 * time.Second)
	fmt.Println("Parent has done its work")
}

func main() {
	fmt.Println("Main has started")
	defer fmt.Println("Main has exited")
	wg.Add(1)
	// go parent() // can be used to check how child continues running if the parent has exited
	go sincereParent() // can be used to check if we want to exit the child if the parent has exited
	wg.Wait()
}
