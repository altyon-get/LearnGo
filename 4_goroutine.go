package main

import (
	"fmt"
	"time"
)

// Function to be executed as a goroutine
func printNums() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second) // Sleep for a second to simulate work
	}
}

func Test3() {
	// Launching the goroutine
	go printNums()

	// Main function doing some other work
	fmt.Println("Main function is running...")

	// Sleep to give time for the goroutine to finish its task
	time.Sleep(6 * time.Second)
	fmt.Println("Main function is done.")
}
