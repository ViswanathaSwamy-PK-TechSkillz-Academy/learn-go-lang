// FileName: main.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Showing the Values Demo")

	// Strings
	fmt.Println("\nShowing Strings")

	fmt.Println("Hello World")
	fmt.Println(`Hello World`)

	fmt.Println("\nInterpreted string :: This is first line. \nThis is second line.\n") // Interpreted string
	fmt.Println(`Raw string :: This is first line. \nThis is second line.`)             // Raw string

	/*
		.\main.go:18:36: newline in string || .\main.go:18:36: syntax error: unexpected newline in argument list; possibly missing comma or )
		.\main.go:19:25: newline in string
	*/
	fmt.Println("This is first line. This is second line.")
	fmt.Println(`
	<Person>
		<FirstName>John</FirstName>
		<LastName>Doe</LastName>
	</Person>
	`)

	// Numbers (int, uint, float32, float64)
	fmt.Println("\nShowing Numbers")
	fmt.Println("Numbers: ", 42)
	fmt.Println("Numbers: ", -42)
	fmt.Println("Numbers: ", 384748542.048575)
	fmt.Println("Numbers: ", 42.0+3.0)
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans
	fmt.Println("\nShowing Booleans")
	fmt.Println("Boolean: ", true)
	fmt.Println("Boolean: ", false)
	fmt.Println("Boolean: ", !true)
	fmt.Println("Boolean: ", !false)
}
