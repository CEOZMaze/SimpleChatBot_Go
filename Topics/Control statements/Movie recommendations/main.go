package main

import "fmt"

func main() {
	var age int
	fmt.Scanf("%d", &age)

	// Code your switch or if...else-if statement here.
	switch {
	case age < 15:
		fmt.Println("Toy Story 4")
	case age > 14 && age < 19:
		fmt.Println("The Matrix")
	case age > 18 && age < 26:
		fmt.Println("John Wick")
	case age > 25 && age < 36:
		fmt.Println("Constantine")
	case age > 35:
		fmt.Println("Speed")

	}
}
