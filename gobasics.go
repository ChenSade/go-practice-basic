package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

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
}

// ------------------------------------------------------------------------

// calculate() returns the sum of the two parameter s
func calculate(value1 string, value2 string) float64 {
	var num1 float64 = convertToFloat(value1)
	var num2 float64 = convertToFloat(value2)

	var sum float64 = math.Round((num1+num2)*100) / 100
	fmt.Printf("the sum of %v and %v is: %v\n\n", num1, num2, sum)
	return sum
}

func convertToFloat(value string) float64 {
	num, err := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if err != nil {
		fmt.Println(err)
		panic("value must be a number")
	}

	return num
}

// ------------------------------------------------------------------------

// Converts a slice of strings to a map object.
func convertToMap(items []string) map[string]float64 {

	result := make(map[string]float64)
	price := 100 / float64(len(items))

	for _, fruit := range items {
		result[fruit] = price
	}
	return result
}

// ------------------------------------------------------------------------

type cartItem struct{
    name string
    price float64
    quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
    var totalVal float64 = 0
    for _, item := range cart {
        totalVal += item.price * float64(item.quantity)
    }
    return totalVal
}

func execCalculateTotal() {
    var cart []cartItem
    var apples = cartItem{"apple", 2.99, 3}
    var oranges = cartItem{"orange", 1.50, 8}
    var bananas = cartItem{"banana", .49, 12}
    cart = append(cart, apples, oranges, bananas)
    total := calculateTotal(cart)
    fmt.Println("\nTotal cart value:", total)
}
// ------------------------------------------------------------------------

func execCalculate2() {
    value1 := "10"
    value2 := "5.5"
    operation := "-"
    result := calculate2(value1, value2, operation)
    fmt.Println("Calculate2 result:", result)
}

// calculate() returns the result of the requested operation.
func calculate2(input1 string, input2 string, operation string) float64 {
    value1 := convertInputToValue(input1)
    value2 := convertInputToValue(input2)

    var result float64

    switch operation {
    case "+":
        result = addValues(value1, value2)
    case "-":
        result = subtractValues(value1, value2)
    case "*":
        result = multiplyValues(value1, value2)
    case "/":
        result = divideValues(value1, value2)
    default:
        fmt.Printf("Invalid operation: %v\n", operation)
        panic("Invalid operation")
    }
	
    return result
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
    if err != nil {
        message := fmt.Sprintf("not a number: %v", input)
        panic(message)
    }
    return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

// ------------------------------------------------------------------------

type cartItem2 struct {
    Name string
    Price float64
    Quantity int
}

func readFromJsonTest() {
    jsonString := 
    `[{"name":"apple","price":2.99,"quantity":3},
    {"name":"orange","price":1.50,"quantity":8},
    {"name":"banana","price":0.49,"quantity":12}]`

    // option 1
    cartItems := getCartFromJson(jsonString)
    fmt.Println("Read from Json 1:")
    for _, item := range cartItems {
        fmt.Printf("Name: %v, Price: %v, Quantity: %v\n", item.Name, item.Price, item.Quantity)
    }

    // option 2
    cartItems = getCartFromJson2(jsonString)
    fmt.Println("Read from Json 2:")
    for _, item := range cartItems {
        fmt.Printf("Name: %v, Price: %v, Quantity: %v\n", item.Name, item.Price, item.Quantity)
    }
}

func getCartFromJson(jsonString string) []cartItem2 {
    cartItems := make([]cartItem2, 0, 10)

    decoder := json.NewDecoder(strings.NewReader(jsonString))
    _, err := decoder.Token()
    if err != nil {
        panic(err)
    }

    var item cartItem2
    for decoder.More() {
        err := decoder.Decode(&item)
        if err != nil {
            panic(err)
        }
        cartItems = append(cartItems, item)
    }
    
    return cartItems
}

func getCartFromJson2(jsonString string) []cartItem2 {
    var cartItems []cartItem2
    err := json.Unmarshal([]byte(jsonString), &cartItems)
    if (err != nil) {
        panic(err)
    }
    return cartItems
}