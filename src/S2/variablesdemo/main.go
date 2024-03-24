// FileName: main.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Showing the Values Demo")

	// Strings
	var msg1 string      // Variable Declaration
	msg1 = "Hello World" // Variable Initialization
	fmt.Println(msg1)

	var msg2 string = "Hello World" // Variable Declaration and Initialization
	fmt.Println(msg2)

	var msg3 = "Hello World" // Type Inference. Initialization based on the value
	fmt.Println(msg3)

	msg4 := "Hello World" // Short Declaration. Type Inference
	fmt.Println(msg4)

	msg4 = "This is first line. \nThis is second line."
	fmt.Println(msg4)

	msg4 = `This is first line. \nThis is second line.`
	fmt.Println(msg4)

	msg4 = `
		<Person>
			<FirstName>John</FirstName>
			<LastName>Doe</LastName>
		</Person>
	`
	fmt.Println(msg4)

	// Numbers
	var n1 = -42
	fmt.Println(n1)

	var n2 uint = 42
	fmt.Println(n2)

	// Error
	n, e := fmt.Println("Hello" + " " + "World")
	fmt.Println("Number of bytes written: ", n)
	fmt.Println("Error: ", e)

}

// // Character
// a, b := 'A', '世'
// fmt.Println(a, string(a))
// fmt.Println(b, string(b))
// fmt.Printf("%c %c\n", a, b)
