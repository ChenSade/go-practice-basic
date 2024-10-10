package main

import (
	"fmt"
)

func printGeneric[T any](t T) string {
	return fmt.Sprintf("%v\n", t)
}

func printSlice[T any](items []T) {
	for i, item := range items {
		fmt.Printf("%d: %s\n", i, printGeneric[T](item))
	}
}

func testPrintSlice() {
	stringSlice := []string { "one", "two", "three", "four"}
	
	// with type inference
	printSlice(stringSlice)
	// without type inference)
	printSlice[string](stringSlice)
}

// ------------------------------------------------------------------------

type Solar struct {
	Name string
	Netto float64
}

type Wind struct {
	Name string
	Netto float64
}

type Yada struct {
	Name string
	Netto float64
}

type Energy interface {
	Solar | Wind
	Cost() float64
}

func (s Solar) Cost() float64 {
	return s.Netto * 0.4
}

func (w Wind) Cost() float64 {
	return w.Netto * 0.6
}
