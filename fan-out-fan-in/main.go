package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. The Conveyor Belts
	rawPhotos := make(chan int)    // Channel for raw photos (ID numbers)
	editedPhotos := make(chan int) // Channel for finished photos

	// 2. THE PRODUCER (The Photographer)
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("📸 Photographer: Took photo #%d\n", i)
			rawPhotos <- i // SEND: Into the raw pipe
		}
		close(rawPhotos) // Done taking photos
	}()

	// 3. THE WORKERS (The Editors - FAN-OUT)
	// We start 3 separate workers (goroutines) to work IN PARALLEL.
	// This is where your mind opens: they all listen to the SAME channel!
	var wg sync.WaitGroup
	for workerID := 1; workerID <= 3; workerID++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for photo := range rawPhotos {
				// RECEIVE: Whichever worker is free grabs the next photo
				fmt.Printf("🎨 Editor %d: Editing photo #%d...\n", id, photo)
				time.Sleep(2 * time.Second) // Editing takes a long time!
				
				// SEND: Into the edited pipe
				editedPhotos <- photo
			}
		}(workerID)
	}

	// 4. THE CLEANUP GOROUTINE
	// This waits for all editors to finish, then closes the final pipe.
	go func() {
		wg.Wait()
		close(editedPhotos)
	}()

	// 5. THE CONSUMER (The Album Creator - FAN-IN)
	fmt.Println("📖 Album Creator: Waiting for finished photos...")
	for finalPhoto := range editedPhotos {
		fmt.Printf("📖 Album Creator: Pasted photo #%d into the book! ✅\n", finalPhoto)
	}

	fmt.Println("🏁 Factory: All work finished. The album is ready!")
}