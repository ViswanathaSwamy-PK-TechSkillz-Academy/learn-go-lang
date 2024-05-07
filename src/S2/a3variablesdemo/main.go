// FileName: main.go

package main

import (
	"fmt"
)

// var empAge int = 25 // This will work. Variable Declaration at the package level
// empName := "John Doe" // This will not work. Short Declaration can't be used at the package level

func main() {
	fmt.Println("Showing the Variables Demo")

	// Show the Zero Values Demo
	fmt.Println("\nShowing Zero Values ...")
	var z1 int
	fmt.Println("Int: ", z1)
	var z2 uint
	fmt.Println("UInt: ", z2)
	var z3 float32
	fmt.Println("Float32: ", z3)
	var z4 float64
	fmt.Println("Float64: ", z4)
	var s1 string
	fmt.Println("String: ", s1, " Length: ", len(s1))
	var b3 bool
	fmt.Println("Boolean: ", b3)

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

	// Multi-line Strings. Escape characters will be honored
	msg4 = "This is first line. \nThis is second line."
	fmt.Println("Message 4: ", msg4)

	// Multi-line Strings using backticks. It will Ignore the formatting. Escape characters will be ignored
	msg4 = `This is first line. \nThis is second line.`
	fmt.Println("Message 4: ", msg4)

	// Multi-line Strings using backticks. It preserves the formatting
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

	var i int = 42
	var f float32 = float32(i)
	fmt.Println("Float32: ", f)

	var f32 float32 = 3.14
	var f64 float64 = float64(f32)
	fmt.Println("Float64: ", f64)

	// var f64_1 float64 = 938493.143847
}
