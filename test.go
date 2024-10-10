package main

import (
	"errors"
	"fmt"
)

func printNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send data into the channel
	}
	close(ch) // Close the channel when done
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func Test() {
	ch := make(chan int)
	go printNumbers(ch)

	for num := range ch { // Read data from the channel
		fmt.Println(num)
	}
	// return

	defer fmt.Println("Deferred call")
	// return
	fmt.Println("Regular call")

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
