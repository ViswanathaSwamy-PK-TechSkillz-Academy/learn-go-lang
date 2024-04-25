package main

import (
	"bankapp/transfers"
	"fmt"
	"time"
)

// main is the start point in golang
func main() {

	time := time.Now()
	transfer := transfers.NewTransfer("1234", time, 12000.7, 44354363)

	fmt.Printf("*** Amount limit validation *** \n")
	// Validate amount
	if transfer.IsAmountLimitAllowed() {
		fmt.Println("Amount is allowed")
	} else {
		fmt.Println("Amount is not allowed")
	}

	fmt.Printf("\n\n *** Date validation: *** \n")
	// Validate date
	if transfer.IsDateAllowed() {
		fmt.Println("Date is allowed")
	} else {
		fmt.Println("Date is not allowed")
	}

}
