package main

import "fmt"

type Car struct {
	brand string
}

func (c Car) drive() {
	fmt.Println(c.brand, "is driving")
}

func Opps() {
	myCar := Car{brand: "Toyota"}
	myCar.drive()
}
