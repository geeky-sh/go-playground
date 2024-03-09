# go-slate
Fiddling with different libraries and concepts in Golang

# Concepts Added
- [slog](https://pkg.go.dev/golang.org/x/exp/slog)
- [mock](https://pkg.go.dev/github.com/stretchr/testify/mock#hdr-Example_Usage)
- [closure](https://gobyexample.com/closures)
- [Ticker](https://gobyexample.com/tickers)
- [Worker pools](https://gobyexample.com/worker-pools)
- [Stateful GoRoutines](https://gobyexample.com/stateful-goroutines)
- [waitgroup](https://gobyexample.com/waitgroups)
- [mutex](https://gobyexample.com/mutexes)
- [Rate Limiting](https://gobyexample.com/rate-limiting)

# Problems (Tools to solve common encountered problems.)
- Build Unix Like `wc` tool. [ref](https://codingchallenges.fyi/challenges/challenge-wc/)
- Build JSON parser [ref](https://codingchallenges.fyi/challenges/challenge-json-parser/)
- Create a custom timeout handler. This is typically used to end any long-running request during API development.

## Context

## Concurrency Patterns
Go allows us to call a function using `go` and it will run as a separate goroutine. This is a very powerful primitive which allows us to easily parallelize parts of our code. However, we we run concurrent programs, we encounter problems that are common to this domain. These are:
#### What are some concurrency primitives provided by Golang
```golang
func add(a, b int) {
    return a + b
}

// invoke function as a go routine
go add(1, 2)

// create channel
ch1 := make(chan int) // sends and receives to an unbuffered channel is blocking


// create channel with buffer
ch2 := make(chan int, 100)

// send messages to a channel
ch1 <- 2

// receive messages from a channel
val := <-ch1

// close channel
v, ok := <-ch1 // ok is false when the channel is closed

// range on channel
for i := range ch1 // receives values from the repeatably until it is closed.
```
#### How to modify shared variables atomically and sequentially.
- Use mutexes [ref](./concepts/mutex/main.go)
- Use channels [ref](./concepts/statefulgoroutines/main.go)
- Use atomic counters [ref](./concepts/atomiccounters/main.go)

#### How to ensure equal distribution of messages among multiple goroutines created to do a specified job.
A common example of encountering this use case is when we have to create a worker pool that can listen to a set of jobs on a channel. This is used when executing DB queries in the application.

In Golang, if the channel is shared between multiple go-routines on the receiving side and when a barrage of messages are sent to the channel, those messages are equally distributed among all the go-routines on the receiving end. This is a very nice feature that is provided by default. [ref](./concepts/workerpools/main.go)
#### How to wait for spawned goroutines to complete
- Use waitgroup [ref](./concepts/waitgroup/main.go)
- Use blocking receive functionality of channel [ref](./concepts/channel/cl.go)

#### Any other useful packages that make it easy to work with concurrent goroutines
- [context](https://pkg.go.dev/context) package. This is used to propagate timeline and deadline signals from one go-routine to another. A common example to use this is to implement a timeout for API requests. [ref](/problems/timeouts/main.go)
- [Timer](https://gobyexample.com/timers) and [Ticker](https://gobyexample.com/tickers). This is used when we want to get a message once in the future or repeatedly. One real-world example of where it can be used is to implement [rate-limiter](./concepts/ratelimiting/main.go)

## Common Concurrency Coding Patterns
#### TODO
