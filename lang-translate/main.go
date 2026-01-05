package main

import (
	"fmt"
	"time"
)

// translator acts like a service that runs in the background
func translator(request chan string, response chan string) {
	// The translator stays in the booth (loop) forever
	for {
		// 📥 STEP 2: Wait for the tourist to say something
		word := <-request

		// Logic: A very simple "dictionary"
		translated := ""
		switch word {
		case "hello":
			translated = "Hola"
		case "bread":
			translated = "Pan"
		case "water":
			translated = "Agua"
		default:
			translated = "I don't know that word"
		}

		// 📤 STEP 3: Send the answer back through the response pipe
		response <- translated
	}
}

func main() {
	// Create the two pipes for communication
	toTranslator := make(chan string)
	fromTranslator := make(chan string)

	// Start the translator worker
	go translator(toTranslator, fromTranslator)

	// 📝 The Tourist's Journey
	wordsToAsk := []string{"hello", "bread", "water"}

	for _, w := range wordsToAsk {
		fmt.Printf("Tourist: How do you say '%s'?\n", w)

		// 1. Send the word to the translator
		toTranslator <- w

		// 2. Wait right here for the answer
		result := <-fromTranslator

		fmt.Printf("Translator says: '%s'\n\n", result)
		
		// Wait a second so we can read the output easily
		time.Sleep(1 * time.Second)
	}

	fmt.Println("All words translated. Closing the booth.")
}