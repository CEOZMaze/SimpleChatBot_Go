package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)

	// Create intSlice containing the a, b and c variables below:
	var intSlice = [3]int{a, b, c}

	fmt.Println(intSlice) // print intSlice here!
}
