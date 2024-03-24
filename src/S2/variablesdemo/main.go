// FileName: main.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Showing the Values Demo")

	// Strings
	fmt.Println("\nShowing Strings ...")
	var msg1 string      // Variable Declaration
	msg1 = "Hello World" // Variable Initialization
	fmt.Println("Message 1: ", msg1)

	var msg2 string = "Hello World" // Variable Declaration and Initialization
	fmt.Println("Message 2: ", msg2)

	var msg3 = "Hello World" // Type Inference. Initialization based on the value
	fmt.Println("Message 3: ", msg3)

	msg4 := "Hello World" // Short Declaration. Type Inference
	fmt.Println("Message 4: ", msg4)

	msg4 = "This is first line. \nThis is second line."
	fmt.Println("Message 4: ", msg4)

	msg4 = `This is first line. \nThis is second line.`
	fmt.Println("Message 4: ", msg4)

	msg4 = `
		<Person>
			<FirstName>John</FirstName>
			<LastName>Doe</LastName>
		</Person>
	`
	fmt.Println("Message 4: ", msg4)

	// Numbers
	fmt.Println("\nShowing Numbers ...")
	var n1 = -42
	fmt.Println("Int: ", n1)

	var n2 uint = 42
	fmt.Println("UInt: ", n2)

	var f1 float32 = 3.14
	fmt.Println("Float32: ", f1)

	var f2 float64 = 938493.143847
	fmt.Println("Float64: ", f2)

	// Boolean
	fmt.Println("\nShowing Boolean ...")
	var b1 bool = true
	fmt.Println("Boolean: ", b1)

	b2 := false
	fmt.Println("Boolean: ", b2)

	// Error
	fmt.Println("\nShowing Error ...")
	n, e := fmt.Println("Message: Hello" + " " + "World")
	fmt.Println("Number of bytes written: ", n)
	fmt.Println("Error: ", e)
}

// // Character
// a, b := 'A', '世'
// fmt.Println(a, string(a))
// fmt.Println(b, string(b))
// fmt.Printf("%c %c\n", a, b)
