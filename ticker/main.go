package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. The News Wire (Channel)
	newsWire := make(chan string)

	// 2. The Reporter (Producer Goroutine)
	// This worker sends a news update every 2 seconds automatically.
	go func() {
		// A Ticker is like an alarm clock that repeats
		ticker := time.NewTicker(2 * time.Second)
		
		headlines := []string{
			"Sun rises in the East",
			"Go is still the best for concurrency",
			"Coffee prices are rising",
			"AI learns to bake pizza",
		}

		for _, msg := range headlines {
			// Wait for the clock to "tick"
			<-ticker.C 
			
			fmt.Println("🎥 Reporter: Breaking news is ready!")
			// SEND: Inward arrow to channel
			newsWire <- msg 
		}

		// Cleanup
		ticker.Stop()
		close(newsWire)
	}()

	// 3. The News Anchor (Consumer / Main Thread)
	fmt.Println("📺 News Anchor: Waiting for updates from the reporter...")

	// This loop waits for the channel to "pour" data out
	for report := range newsWire {
		// RECEIVE: The report variable catches data from newsWire
		fmt.Printf("📺 News Anchor: 'Our top story tonight: %s'\n\n", report)
	}

	fmt.Println("🏁 News Broadcast: Signing off. Goodnight!")
}