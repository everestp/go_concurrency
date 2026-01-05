package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		//print  a got data message
		i := <-ch
		fmt.Println("Got", i, "Form the channel")

		//simulate doing a lot of work
		time.Sleep(8 * time.Second)

	}
}


func main(){
	ch := make(chan int ,10)
	go listenToChan(ch)
	for i := 0 ; i <=100 ; i++{
		fmt.Println("sending ",i ,"to channel...")
		ch <-i
		fmt.Println("Sent ",i,"to channel")

	}

	fmt.Println("done")
	close(ch)

}