package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. The Conveyor Belt (Channel)
	// This belt carries strings (package names)
	belt := make(chan string)

	// 2. The Packer (Producer Goroutine)
	go func() {
		packages := []string{"iPhone", "Laptop", "Monitor", "Keyboard"}

		for _, item := range packages {
			fmt.Printf("📦 Packer: Boxed up a %s. Putting it on the belt...\n", item)
			
			// SEND: This "pokes" the main thread
			belt <- item 

			// Wait a bit before packing the next one
			time.Sleep(1 * time.Second) 
		}

		// IMPORTANT: Stop the belt so the Truck knows to leave
		fmt.Println("🛑 Packer: No more packages. Closing the belt.")
		close(belt)
	}()

	// 3. The Shipping Truck (Consumer / Main Thread)
	fmt.Println("🚚 Truck: Waiting at the dock...")

	// This loop keeps the program alive and running!
	for p := range belt {
		// As long as data is sent, the program continues here:
		fmt.Printf("🚚 Truck: Received %s! Scanning and loading...\n", p)
		
		// Simulate the truck taking time to scan/load the item
		time.Sleep(500 * time.Millisecond)
	}

	// 4. The Exit
	// This only happens AFTER the channel is closed.
	fmt.Println("🏁 Truck: Belt is empty and closed. Driving away. Goodbye!")
}