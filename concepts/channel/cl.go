package main

import (
	"fmt"
	"net/http"
	"time"
)

func greetMsg(ch chan string, name string) {
	fmt.Println("Is this printed at all?")
	ch <- fmt.Sprintf("Hello %s! Welcoome to the world of channels", name)
	time.Sleep(4 * time.Second)
	fmt.Println("end of goroutine")
	http.Get("https://en5azdeb65cgy.x.pipedream.net")
}

func channelIntro() {
	ch := make(chan string)

	go greetMsg(ch, "Aash")

	msg := <-ch
	fmt.Println(msg)
	fmt.Println("This is the almost the end of main function")
	time.Sleep(3 * time.Second)
	fmt.Println("This is  the end of main function")
}

func channelBuffering() {
	ch := make(chan string, 1)

	ch <- "message1"
	ch <- "message2"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func main() {
	channelBuffering()
}
