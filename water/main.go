package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. The Pipe (Channel)
	// This carries "liters" of water (integers)
	waterPipe := make(chan int)

	// 2. The Pump (Producer Goroutine)
	// We hire ONE worker to pump 5 liters of water
	go func() {
		for liter := 1; liter <= 5; liter++ {
			fmt.Printf("🚰 Pump: Pumping Liter #%d into the pipe...\n", liter)
			
			// SEND: Pushing data into the pipe
			waterPipe <- liter 

			// The pump needs time to reset between liters
			time.Sleep(800 * time.Millisecond) 
		}

		// When finished, we must "turn off the valve"
		fmt.Println("🚫 Pump: Tank empty. Closing pipe.")
		close(waterPipe)
	}()

	// 3. The Filter (Consumer / Main Thread)
	fmt.Println("🧪 Filter: Ready to process...")

	// This loop runs as long as the Pump is sending water
	for drop := range waterPipe {
		fmt.Printf("🧪 Filter: Receiving Liter #%d...\n", drop)
		
		// Simulate "Processing" the data
		fmt.Printf("✨ Filter: Liter #%d is now CLEAN! ✅\n", drop)
		
		// The filter is very fast (200ms)
		time.Sleep(200 * time.Millisecond)
	}

	// 4. The Exit
	fmt.Println("🏁 System: Pipe closed. Filtration complete.")
}