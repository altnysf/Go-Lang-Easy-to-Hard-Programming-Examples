package main

// Area of Rectangle : LongSide (a) - ShortSide(b) : a x b

// Perimeter of Rectangle : 2 x (a + b) --> ( 2 x a ) + ( 2 x b) --> 2a + 2b

import (
	"fmt"
)

func main() {

	var a, b float64
	fmt.Print("LongSide of Rectangle :")
	fmt.Scanln(&a)

	fmt.Print("ShortSide of Rectangle :")
	fmt.Scanln(&b)

	area := a * b
	fmt.Printf("Area of Rectangle is %v \n", area)

	perimeter := 2 * (a + b)
	fmt.Printf("Perimeter of Rectangle is %v \n", perimeter)

}
