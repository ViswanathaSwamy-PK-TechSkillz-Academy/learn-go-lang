// FileName: main.go

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("\nShowing the executable path.")

	exePath, err := os.Executable()

	if err != nil {
		fmt.Println("\nFailed to get executable path: ", err)
	} else {
		fmt.Println("\nExecutable Path : ", exePath)
	}

}
