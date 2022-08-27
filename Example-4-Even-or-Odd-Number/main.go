package main

// This program that prints whether the number typed by the user is a even or odd number.

import "fmt"

func main() {
	var number int
	fmt.Print("Write a number : ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("The Number is Even :", number)
	} else if number%2 != 0 {
		fmt.Println("The Number is Odd :", number)
	} else {
		fmt.Println("There is an Error..!!")
	}
}
