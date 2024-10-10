package main

import (
	"fmt"
	"rsc.io/quote/v4"
)

func main() {
	fmt.Println("Hi Everyone!!!")
	fmt.Println(quote.Go())

	// calculate
	calculate("10", "5.5")

	// Collections
	slice := []string{"apple", "banana", "orange"}
	result := convertToMap(slice)
	fmt.Printf("%v", result)

    // program flow
    execCalculateTotal()

    // calculate-2
    execCalculate2()

    // read from JSON 
    readFromJsonTest()

	testPrintSlice()
}