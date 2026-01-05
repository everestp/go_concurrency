package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1. One "Result Pipe" (Channel)
	// Both workers will try to pour their answer into this same pipe.
	resultPipe := make(chan string)

	// 2. Start Worker A (Google)
	go func() {
		delay := rand.Intn(5) // Random speed 0-5 seconds
		time.Sleep(time.Duration(delay) * time.Second)
		resultPipe <- "Google finished first!" // SEND
	}()

	// 3. Start Worker B (Bing)
	go func() {
		delay := rand.Intn(5) // Random speed 0-5 seconds
		time.Sleep(time.Duration(delay) * time.Second)
		resultPipe <- "Bing finished first!" // SEND
	}()

	// 4. THE MIND-BLOWER: THE FIRST CATCH
	// We only listen to the channel ONE TIME.
	// As soon as the FASTEST worker sends their message, 
	// the main thread takes it and moves to the next line immediately.
	fmt.Println("🚀 Waiting for the fastest result...")
	
	winner := <-resultPipe // RECEIVE: Only catches the first one to arrive!

	fmt.Println("🏆 Winner:", winner)
	fmt.Println("🏁 Closing the program. We don't need to wait for the slow one!")
}