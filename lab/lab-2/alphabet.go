package main

import (
	"fmt"
	"time"
	)

func printAlphabet() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf(string(i));
	}
	fmt.Println("\nDone printing the alphabet")
}

func main() {
	go printAlphabet()
	fmt.Printf("\nWill sleep it out...")
	for i:= 1; i <= 10; i++ {
		fmt.Printf(string('z'))
		time.Sleep(1*time.Second)
	}
}