package main

// Write a program that performs addition, subtraction, multiplication, division calculations of 2 numbers written by the user.

import "fmt"

func main() {
	var number1, number2 float64

	fmt.Print("First Number : ")
	fmt.Scanln(&number1)

	fmt.Println("-----------------------------")

	fmt.Print("Second Number : ")
	fmt.Scanln(&number2)

	// Addition

	result := number1 + number2
	fmt.Printf("\nSum of Number1 and Number2 : %v\n", result)

	// Subtraction

	if number1 > number2 {
		result := number1 - number2
		fmt.Printf("\nSubtraction of Number1 and Number2 : %v\n", result)
	} else {
		result := number2 - number1
		fmt.Printf("\nSubtraction of Number2 and Number1 : %v\n", result)
	}

	// Multiplication

	result = number1 * number2
	fmt.Printf("\nMultiplication of Number1 and Number2 : %v\n", result)

	// Division

	if number1 > number2 {
		result := number1 / number2
		fmt.Printf("\nDivision of Number1 and Number2 : %v\n", result)
	} else {
		result := number2 / number1
		fmt.Printf("\nDivision of Number2 and Number1 : %v\n", result)
	}

}
