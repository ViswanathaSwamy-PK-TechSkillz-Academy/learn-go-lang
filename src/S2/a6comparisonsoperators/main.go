package main

import (
	"fmt"
)

func main() {
	fmt.Println("Showing Comparision Operators Demo")

	n1, n2 := 5, 2
	fmt.Println("\nComparison Operators")
	fmt.Println("Equality: ", n1 == n2)
	fmt.Println("Inequality: ", n1 != n2)
	fmt.Println("Less Than: ", n1 < n2)
	fmt.Println("Less Than or Equal: ", n1 <= n2)
	fmt.Println("Greater Than: ", n1 > n2)
	fmt.Println("Greater Than or Equal: ", n1 >= n2)

	// String Comparison
	s1, s2 := "Hello", "World"
	fmt.Println("\nString Comparison")
	fmt.Println("Equality: ", s1 == s2)
	fmt.Println("Inequality: ", s1 != s2)
}
