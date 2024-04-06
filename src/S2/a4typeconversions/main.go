// FileName: main.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Showing Type Conversions Demo")

	var i int = 10
	fmt.Println("Int: ", i)
	var f float64 = 10
	fmt.Println("Float64: ", f)

	// Type Conversion
	var i2 int = 10
	fmt.Println("Int: ", i2)

	// cannot use i2 (variable of type int) as float64 value in variable declaration
	// var f22 float64 = i2

	// Go does not support implicit type conversion
	var f2 float64 = float64(i2)
	fmt.Println("Int to Float64: ", f2)

	sal := 12345.678
	fmt.Println("Salary: ", sal)

	// var salInt int = sal // cannot use sal (variable of type float64) as int value in variable declaration
	var salInt int = int(sal)
	fmt.Println("Salary to Int: ", salInt)

	var i8 int8 = 10
	fmt.Println("Int8: ", i8)

	// var i16 int16 = i8 // cannot use i8 (variable of type int8) as int16 value in variable declaration
	var i16 int16 = int16(i8)
	fmt.Println("Int8 to Int16: ", i16)
}
