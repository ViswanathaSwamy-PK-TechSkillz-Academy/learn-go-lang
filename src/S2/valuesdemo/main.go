// FileName: main.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Showing the Values Demo")

	// Character
	fmt.Println('A', rune('A'), string('A'))
	fmt.Println('世', rune('世'), string('世'))

	// Strings
	fmt.Println("Hello World")
	fmt.Println(`Hello World`)

	fmt.Println("This is first line. \nThis is second line.")
	fmt.Println(`This is first line. \nThis is second line.`)

	/*
		.\main.go:18:36: newline in string || .\main.go:18:36: syntax error: unexpected newline in argument list; possibly missing comma or )
		.\main.go:19:25: newline in string
	*/
	fmt.Println("This is first line. This is second line.")
	fmt.Println(`This is first line. 
	This is second line.`)

}
