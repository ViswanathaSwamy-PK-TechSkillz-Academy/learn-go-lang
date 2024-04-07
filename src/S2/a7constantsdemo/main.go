package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nShowing Constants Demo")

	var i int = 10
	fmt.Println("Int: ", i)
	var f float64 = 10
	fmt.Println("Float64: ", f)

	// untyped constant OR Implicitly typed constant
	const a = 42
	fmt.Println("Untyped Constant: ", a)

	// One constant can be assigned to another constant of same type
	const a1 int = a
	fmt.Println("Typed int Constant: ", a1)

	const a2 float32 = a
	fmt.Println("Typed float32 Constant: ", a2)

	// typed constant OR Explicitly typed constant
	const f32 float32 = 3.14
	fmt.Println("Typed float32 Constant: ", f32)

	// cannot use f32 (constant 3 of type float32) as float64 value in constant declaration
	// const f64 float64 = f32
	const f64 float64 = float64(f32)
	fmt.Println("Typed float64 Constant: ", f64)

	// Constants can be used in expressions
	const cccc = 2 * 5
	fmt.Println("Constant Expression 2 * 5 : ", cccc)

	const dddd = "Hello " + "Gophers"
	fmt.Println("Constant Expression 'Hello ' + 'Gophers' : ", dddd)
}
