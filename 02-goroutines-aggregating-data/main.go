package main

import (
	"fmt"
	"sync"
	"time"
)

// the main function
func main() {
	start := time.Now() // start time
	userName := FetchUser() // get the user name
	respch := make(chan any, 2) // create a channel to store the response
	wg := &sync.WaitGroup{} // create a wait group

	wg.Add(2) // add two goroutines to the wait group


	go FetchUserLikes(userName, respch, wg) // fetch user's likes
	go FetchUserMatch(userName, respch, wg) // fetch user's matches

	wg.Wait() // block until 2 wg.Done()
	close(respch) // close the channel

	for resp := range respch {
		fmt.Println("resp: ", resp)
	}

	fmt.Println("took:", time.Since(start)) // print the elapsed time
}

// fetchUser fetches user's name
func FetchUser() string {
	time.Sleep(time.Millisecond * 100) // sleep for 100 ms

	return "BOB"
}

// fetchUserLikes is a goroutine that fetches the user's likes
func FetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150) // sleep for 150ms

	respch <- 11 // send the response to the channel
	wg.Done()
}

// fetchUserMatch is a goroutine that fetches the user's matches
func FetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100) // sleep for 100 ms

	respch <- "ANNA" // send the response to the channel
	wg.Done()
}