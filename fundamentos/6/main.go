package main

import "fmt"

// Slice é um pedaço de um array

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[4:]), cap(s[4:]), s[4:])

	s = append(s, 110)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s)

}
