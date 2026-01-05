package main

import "testing"
func Test_updateMessage(t *testing.T){
	msg = "Hello world"

	wg.Add(1)
	go updateMessage("1")
	go updateMessage("Goodbye ,Cruel world")
	wg.Wait()

	if msg != "Goodbye ,Cruel world"{
		t.Error("incorrect value in msg")
	}
}