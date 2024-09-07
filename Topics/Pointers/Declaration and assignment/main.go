package main

import "fmt"

func main() {
	var p = new(string) // Create a pointer variable of the 'string' type
	*p = "my string"    // Assign "my string" to the variable at the address held by the pointer

	fmt.Println(*p) // print the variable referenced by your pointer variable here!
}
