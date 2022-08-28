package main

// This program that prints numbers between 2 numbers entered by the user

import "fmt"

func main() {

	var num1, num2 int
	fmt.Print("Number 1 : ")
	fmt.Scanln(&num1)
	fmt.Print("Number 2 : ")
	fmt.Scanln(&num2)

	for i := num1; i <= num2; i++ {
		fmt.Println(i)
	}

}
