package main

import (
	"fmt"
	"time"
)

// 1. Define the item being made
type Pizza struct {
	ID int
}

func main() {
	// 2. Create the "Conveyor Belt" (Channel)
	// It only carries 'Pizza' objects.
	conveyorBelt := make(chan Pizza)

	// 3. Start the Chef (The Producer)
	// 'go' starts this function in the background.
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Chef: Making pizza #%d...\n", i)
			time.Sleep(1 * time.Second) // Simulate cooking time

			// SEND: Put the pizza on the belt
			conveyorBelt <- Pizza{ID: i}
		}
		// When done, stop the belt so the customer knows no more are coming.
		close(conveyorBelt)
	}()

	// 4. The Customer (The Consumer)
	// This loop waits for every item that comes through the belt.
	fmt.Println("Customer: Waiting for pizzas...")
	
	for pizza := range conveyorBelt {
		// RECEIVE: Grab the pizza from the belt
		fmt.Printf("Customer: Received and eating pizza #%d! 🍕\n", pizza.ID)
	}

	fmt.Println("Shop closed.")
}