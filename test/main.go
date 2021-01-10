package main

import (
	"fmt"
)

type Order struct {
	ordId      int
	customerId int
	callback   func()
}

func main() {
	var i interface{}
	i = Order{
		ordId:      456,
		customerId: 56,
	}
	value, ok := i.(Order)
	if !ok {
		fmt.Println("It's not ok for type Order")
		return
	}
	fmt.Println("The value is ", value)
}
