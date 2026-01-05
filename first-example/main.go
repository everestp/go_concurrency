package main

import (
	"fmt"
	"sync"
	"time"
)

func printSomething(s string , wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println(s)
}
 func main(){

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().



 var wg sync.WaitGroup
	words := []string{
		"alphha",
		"beta",
		"gama",
		"pi",
		"theta",
	}
	wg.Add(len(words))

	for i , x := range words{
		go printSomething(fmt.Sprintf("%d: %s", i,x),&wg)
	}
	wg.Wait()
	
	 printSomething("This is thr first thing to be printed",&wg)
	time.Sleep(30 * time.Microsecond)
	printSomething("This is the second thig to print",&wg)
 }