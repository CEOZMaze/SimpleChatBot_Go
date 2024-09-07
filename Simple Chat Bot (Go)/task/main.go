package main

import "fmt"

func main() {
	fmt.Println("Hello! My name is Aid.")
	fmt.Println("I was created in 2023.")
	fmt.Println("Please, remind me of your name.")

	var name string
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")

	var rem3, rem5, rem7, age int
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105

	fmt.Printf("Your age is %d; that's a good time to start programming!\n", age)
	fmt.Println("Now I will prove to you that I can count to any number you want.")
	count := 0

	fmt.Scan(&count)

	for i := 0; i <= count; i++ {
		fmt.Printf("%d !\n", i)
	}
	// read a number and count to it here
	printQuestions()
	evaluateAnswer()

}

func printQuestions() {
	fmt.Print("Let's test your programming knowledge.\nWhy do we use methods?\n")
	questions := []string{"1. To repeat a statement multiple times.", "2. To decompose a program into several small subroutines.", "3. To determine the execution time of a program.", "4. To interrupt the execution of a program"}
	for _, question := range questions {
		fmt.Println(question)
	}
}

func getAnswer() int {
	answer := 0
	fmt.Scanln(&answer)
	return answer
}

func evaluateAnswer() {
	for {
		if getAnswer() != 2 {
			fmt.Println("Please, try again.")
		} else {
			fmt.Println("Congratulations, have a nice day!")
			break
		}
	}

}
