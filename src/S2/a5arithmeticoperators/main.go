package main

import (
	"fmt"
)

func main() {

	fmt.Println("Showing Arithmetic Operators Demo")

	n1, n2 := 5, 2
	fmt.Println("Addition: ", n1+n2)
	fmt.Println("Substraction: ", n1-n2)
	fmt.Println("Multiplication: ", n1*n2)
	fmt.Println("Division: ", n1/n2) // Integer Division
	fmt.Println("Float Division: ", float64(n1)/float64(n2))
	fmt.Println("Modulus: ", n1%n2) // Modulus

	n3, n4 := 25.68, 28.912
	fmt.Println("\nAddition: ", n3+n4)
	fmt.Println("Substraction: ", n3-n4)
	fmt.Println("Multiplication: ", n3*n4)
	fmt.Println("Division: ", n3/n4)
	// Not allowed -> invalid operation: operator % not defined on n3 (variable of type float64)
	// fmt.Println("Modulus: ", n3%n4)
	fmt.Println("Integer Modulus: ", int(n3)%int(n4))
}
