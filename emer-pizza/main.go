package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Our communication channels
	pizzaBelt := make(chan string)
	quitChannel := make(chan bool)

	// 2. THE CHEF (Producer)
	go func() {
		for i := 1; i <= 5; i++ {
			// Simulate a random cooking time (0 to 4 seconds)
			time.Sleep(time.Duration(i) * time.Second)
			
			// SEND: Try to put pizza on belt, but also watch for the quit signal
			select {
			case pizzaBelt <- fmt.Sprintf("Pizza #%d", i):
				// Successfully sent!
			case <-quitChannel:
				fmt.Println("👨‍🍳 Chef: I'm quitting! Throwing away the dough.")
				return // Exit the goroutine
			}
		}
	}()

	// 3. THE MANAGER (Main Thread)
	fmt.Println("🚀 Shop is Open! Waiting for pizzas...")

	for {
		// THE MIND-BLOWER: THE SELECT STATEMENT
		// It waits for whichever case happens FIRST.
		select {
		
		case pizza := <-pizzaBelt:
			// CASE 1: A pizza arrived!
			fmt.Printf("🍕 Manager: Received %s. Selling it now!\n", pizza)

		case <-time.After(3 * time.Second):
			// CASE 2: The "Safety Timer." 
			// If no pizza arrives for 3 seconds, this case triggers.
			fmt.Println("⏰ Manager: Chef is too slow! Closing shop due to inactivity.")
			quitChannel <- true // Tell the chef to stop
			return              // End the program

		case <-time.Tick(10 * time.Second):
			// CASE 3: A "Background Task" 
			// This happens every 10 seconds regardless of other work.
			fmt.Println("🧹 Manager: Doing a quick floor cleaning...")
		}
	}
}