// FileName: main.go

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Showing the executable path.")

	exePath, err := os.Executable()

	if err != nil {
		fmt.Println("Failed to get executable path: ", err)
	} else {
		fmt.Println("Executable Path : ", exePath)
	}

}
