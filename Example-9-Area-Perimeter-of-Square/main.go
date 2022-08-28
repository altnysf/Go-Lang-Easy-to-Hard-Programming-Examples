package main

// Area of Square : a²

// Perimeter of Square : 4a

import (
	"fmt"
)

func main() {

	var a float64
	fmt.Print("Side of Square :")
	fmt.Scanln(&a)

	area := a * a // we can solve this question with math.Pow(a, 2)
	fmt.Printf("Area of Square is %v \n", area)

	perimeter := 4 * a
	fmt.Printf("Perimeter of Square is %v \n", perimeter)

}
