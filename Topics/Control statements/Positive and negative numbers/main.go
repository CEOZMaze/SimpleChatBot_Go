package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d", &number)

	// Write your code here.
	if number > 0 {
		fmt.Println("Positive!")
	} else if number == 0 {
		fmt.Println("Zero!")
	} else if number < 0 {
		fmt.Println("Negative!")
	}
}
