package main

import (
	"fmt"
	"math/rand"
	// import the rand package here.
)

// DO NOT change the code within the main function!
func main() {
	var num int64
	fmt.Scanf("%d", &num)

	r := rand.New(rand.NewSource(num))
	fmt.Println(r.Intn(50))
}
