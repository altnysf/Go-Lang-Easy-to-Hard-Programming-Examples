package main

// This program prints the text entered by the user.

import "fmt"

func main() {
	var username string
	fmt.Scanln(username)

	fmt.Println("Hello ", username)
}
