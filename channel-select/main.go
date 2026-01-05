package main

import (
	"fmt"
	"time"
)

// server1 is a slow producer (6 seconds per message)
func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		// SEND: Inward arrow pushes data into the pipe
		ch <- "This is from server 1"
	}
}

// server2 is a fast producer (3 seconds per message)
func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		// SEND: Inward arrow pushes data into the pipe
		ch <- "This is from server 2 "
	}
}

func main() {
	fmt.Println("Select with channel")
	fmt.Println("-------------------")

	// 1. Create two separate communication pipes
	channel1 := make(chan string)
	channel2 := make(chan string)

	// 2. Launch both servers in the background
	go server1(channel1)
	go server2(channel2)

	// 3. The "Infinite Listener"
	for {
		// THE SELECT BLOCK: 
		// This is like a race. The code pauses here and waits for ANY case to be ready.
		// As soon as one server sends data, that case "wins" and runs.
		select {

		case s1 := <-channel1: // RECEIVE: Outward arrow catches server1
			fmt.Println("case one", s1)

		case s2 := <-channel2: // RECEIVE: Outward arrow catches server2
			fmt.Println("case two", s2)

		case s3 := <-channel1: // Duplicate case: if server1 is ready, Go might pick s1 or s3 randomly
			fmt.Println("case three", s3)

		case s4 := <-channel2: // Duplicate case: if server2 is ready, Go might pick s2 or s4 randomly
			fmt.Println("case four", s4)

		// Note on Default:
		// If you uncomment 'default', the loop will never wait. It will just 
		// print 'default' thousands of times per second until a server is ready.
		// Keeping it commented out is better here because it lets the CPU "rest" 
		// while waiting for the servers.
		
		}
	}
}