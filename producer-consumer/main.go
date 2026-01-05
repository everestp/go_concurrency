package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

/////////////////////////////////////////////////////
// GLOBAL CONFIGURATION & STATE
/////////////////////////////////////////////////////

const NumberOfPizzas = 10

// Global counters (Note: In high-concurrency apps, use sync/atomic or mutexes for these)
var pizzasMade int
var pizzasFailed int
var totalPizzas int

/////////////////////////////////////////////////////
// DATA STRUCTURES
/////////////////////////////////////////////////////

type Producer struct {
	data chan PizzaOrder
	quit chan chan error // A channel that receives a channel (for two-way shutdown signaling)
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

/////////////////////////////////////////////////////
// PRODUCER METHODS
/////////////////////////////////////////////////////

// Close sends a signal to the producer to stop and waits for a response
func (p *Producer) Close() error {
	ch := make(chan error) // Create a temporary "callback" channel
	p.quit <- ch           // Send the callback channel to the producer
	return <-ch            // Block here until the producer sends a value back
}

/////////////////////////////////////////////////////
// BUSINESS LOGIC
/////////////////////////////////////////////////////

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber > NumberOfPizzas {
		return &PizzaOrder{pizzaNumber: pizzaNumber}
	}

	delay := rand.Intn(5) + 1
	fmt.Printf("📥 Received order #%d\n", pizzaNumber)

	rnd := rand.Intn(12) + 1
	message := ""
	success := false

	totalPizzas++

	if rnd < 5 {
		pizzasFailed++
	} else {
		pizzasMade++
	}

	fmt.Printf("🍕 Making pizza #%d (will take %d seconds)\n", pizzaNumber, delay)

	// Simulate the "work" time
	time.Sleep(time.Duration(delay) * time.Second)

	if rnd <= 2 {
		message = fmt.Sprintf("❌ Ran out of ingredients for pizza #%d", pizzaNumber)
	} else if rnd <= 4 {
		message = fmt.Sprintf("🔥 Cook quit while making pizza #%d", pizzaNumber)
	} else {
		success = true
		message = fmt.Sprintf("✅ Pizza #%d is ready!", pizzaNumber)
	}

	// FIX: Create the struct first, then return the pointer to it
	order := PizzaOrder{
		pizzaNumber: pizzaNumber,
		message:     message,
		success:     success,
	}

	return &order
}

/////////////////////////////////////////////////////
// PRODUCER GOROUTINE
/////////////////////////////////////////////////////

func pizzaria(pizzaMaker *Producer) {
	var i int

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
		}

		// The 'select' statement allows us to handle multiple channel operations
		select {
		// Case 1: Send the pizza data to the consumer
		case pizzaMaker.data <- *currentPizza:
			// Success: the consumer received the pizza

		// Case 2: We received a shutdown signal
		case quitChan := <-pizzaMaker.quit:
			close(pizzaMaker.data) // Always close your data channels when finished
			quitChan <- nil        // Tell the 'Close' method we are done cleaning up
			return                 // Exit the goroutine
		}
	}
}

/////////////////////////////////////////////////////
// MAIN FUNCTION
/////////////////////////////////////////////////////

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Cyan("🍕 The Pizzeria is now OPEN for business!")
	color.Cyan("=======================================")

	// Initialize the producer with channels
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// Start the worker in the background (Goroutine)
	go pizzaria(pizzaJob)

	// CONSUMER: Range over the channel until it is closed
	for pizza := range pizzaJob.data {
		if pizza.pizzaNumber > NumberOfPizzas {
			if pizza.success {
				color.Green(pizza.message)
			} else {
				color.Red(pizza.message)
			}
			
			// Trigger the graceful shutdown
			_ = pizzaJob.Close()
			break 
		}

		color.Yellow(pizza.message)
	}

	// Final Report
	color.Cyan("\n📊 PIZZA SUMMARY")
	color.Cyan("----------------")
	color.Cyan("Total pizzas attempted: %d", totalPizzas)
	color.Cyan("Pizzas made successfully: %d", pizzasMade)
	color.Cyan("Pizzas failed: %d", pizzasFailed)
}