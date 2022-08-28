package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Print("Number 1 : ")
	fmt.Scanln(&num1)
	fmt.Print("Number 2 : ")
	fmt.Scanln(&num2)

	for i := num1; i <= num2; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is an Even Number \n", i)
		} else {
			fmt.Printf("%v is an Odd Number \n", i)
		}
	}

}
