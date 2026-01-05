package main

import (
	"fmt"
	"time"
)

// 1. The Data Object
type Transaction struct {
	ID     int
	Amount int
}

func main() {
	// 2. The Communication Line (Channel)
	// This belt carries Transaction structs
	bankRequests := make(chan Transaction)

	// 3. The ATM (Producer Goroutine)
	// One worker sending 3 different transactions
	go func() {
		amounts := []int{100, 500, -200} // Deposit, Deposit, Withdrawal

		for i, amt := range amounts {
			t := Transaction{ID: i + 1, Amount: amt}
			fmt.Printf("🏧 ATM: Sending transaction #%d ($%d)...\n", t.ID, t.Amount)
			
			// SEND: Pushing the struct into the channel
			bankRequests <- t 

			// Wait for the customer to type on the screen
			time.Sleep(1 * time.Second) 
		}

		// Close the line when the ATM session ends
		fmt.Println("🛑 ATM: Session ended. Closing connection.")
		close(bankRequests)
	}()

	// 4. The Bank Server (Consumer / Main Thread)
	balance := 1000
	fmt.Printf("🏦 Server: Starting Balance: $%d\n", balance)

	// The program stays alive here as long as the ATM is sending data
	for req := range bankRequests {
		fmt.Printf("🏦 Server: Processing #%d...\n", req.ID)
		
		// Logic: Update the balance
		balance += req.Amount
		fmt.Printf("🏦 Server: Transaction #%d complete. New Balance: $%d\n", req.ID, balance)
	}

	// 5. Final Result
	fmt.Printf("🏁 Final Bank Balance: $%d\n", balance)
} 

