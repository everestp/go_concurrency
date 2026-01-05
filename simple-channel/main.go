package main

import (
	"fmt"
	"strings"
	"time"
)

// shout is a "Worker Service."
// It takes data from the 'ping' channel, processes it, and sends it back to 'pong'.
func shout(ping chan string, pong chan string) {
	// This infinite loop keeps the worker alive as long as the program runs.
	for {
		// STEP 2 & 6: The worker "blocks" (waits) here until main sends a ping.
		s := <-ping

		// If the channel is closed, 's' will be an empty string. 
		// In a real app, we would check if the channel is closed here.

		// STEP 7: Transform the data and send it back.
		// fmt.Sprintf creates the "shouted" version of the string.
		pong <- fmt.Sprintf("%s !!", strings.ToUpper(s))
	}
}

func main() {
	// 1. Initialize two unbuffered channels.
	// Unbuffered means the sender MUST wait for a receiver.
	ping := make(chan string)
	pong := make(chan string)

	// 2. Launch the worker goroutine.
	// This starts 'shout' in the background. It immediately goes to its loop and waits.
	go shout(ping, pong)

	// Small delay to ensure the goroutine has started (optional but safe).
	time.Sleep(10 * time.Millisecond)

	// NOTE: Changed Sprintf to Println so you can actually see the instruction.
	fmt.Println("Type something and press enter (or 'q' to quit)")

	// 3. The Interactive Loop (The Consumer/UI thread)
	for {
		fmt.Printf("-> ")

		// 4. Capture user input from the terminal.
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		// 5. Check the exit condition.
		if strings.ToLower(userInput) == "q" {
			break
		}

		// STEP 5: SEND phase. 
		// main sends the string to the worker. main blocks here until 'shout' receives it.
		ping <- userInput

		// STEP 8: RECEIVE phase.
		// main blocks here until 'shout' finishes processing and sends the result to 'pong'.
		response := <-pong

		// 9. Display the result to the user.
		fmt.Println("Response:", response)
	}

	// 10. Cleanup Phase
	// Once the loop breaks, we close the channels to free up resources.
	fmt.Println("All done! Closing the channels...")
	close(ping)
	close(pong)
}