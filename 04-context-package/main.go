package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// This example shows how to use a context to cancel a long-running operation.
func main() {
	start := time.Now() // start time
	ctx := context.Background() // use the background context
	userID := 10
	val, err := fetchUserDataByID(ctx, userID) // fetch user data by ID
	if err != nil {
		log.Fatal(err) // handle error
	}

	fmt.Println("result: ", val) // print the result
	fmt.Println("took: ", time.Since(start)) // print the elapsed time

}

// response is the type returned by the fetchUserDataByID function.
type Response struct {
	value int // the value of the user from the third-party service
	err   error // error encountered while fetching the user data from the third-party service
}

// fetchUserDataByID fetches the user data by ID from the third-party service.
func fetchUserDataByID(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200) // set a timeout for the context
	defer cancel() // cancel the context when we're done
	respch := make(chan Response) // create a channel to receive the response

	go func() { // create a goroutine to fetch the user data
		val, err := fetchThirdPartyStuffWichCanBeSlow() // fetch the user data from the third-party service
		respch <- Response{ // send the response to the response channel
			value: val, // the value of the user from the third-party service
			err:   err, // error encountered while fetching the user data from the third-party service
		}
	}()

	for { // wait for the goroutine to finish
		select { // wait for the response channel to receive a value
		case <-ctx.Done(): // if the context is canceled, break out of the select
			return 0, fmt.Errorf("fectching data from third party took to long")
		case resp := <-respch: // if the response channel receives a value, break out of the select
			return resp.value, resp.err // return the value of the user from the third-party service
		}
	}

}

// fetchThirdPartyStuffWichCanBeSlow take up to 500 Milliseconds to fetch the user data from the third-party service.
func fetchThirdPartyStuffWichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 500)
	return 666, nil // return the devil
}
