package main

// This program that prints whether the number typed by the user is a positive or negative number.

import "fmt"

func main() {
	var number int
	fmt.Print("Write a number : ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("The Number is Negative :", number)
	} else if number > 0 {
		fmt.Println("The Number is Positive :", number)
	} else {
		fmt.Println("The Number is Equal to 0 (Zero)")
	}
}
