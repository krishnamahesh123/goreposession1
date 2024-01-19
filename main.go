package main

import (
	"assignment1/addition"
	"assignment1/division"
	"assignment1/multiplication"
	"assignment1/subtraction"
	"fmt"
)

func main() {
	fmt.Println("Below is the calculator that performs Addition, Subtraction, Multiplication and Division")
	res1 := addition.Add(10, 20)
	fmt.Println(res1)
	res2 := subtraction.Sub(10, 20)
	fmt.Println(res2)
	res3 := multiplication.Multiply(10, 20)
	fmt.Println(res3)
	res4, res5 := division.Divide(10, 20)
	fmt.Println(res4)
	fmt.Println(res5)

}
