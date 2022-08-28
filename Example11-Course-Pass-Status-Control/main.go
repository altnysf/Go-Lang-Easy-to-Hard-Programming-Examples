package main

import "fmt"

// "Course Pass Status" Control Program

// RULES - Algorithm
/*
	1 - The student will take 2 exams.
	2 - The average of 2 exams will be calculated
	3 - If the exams average is less than 50, "FAIL" will be displayed, if it is greater than 50, "PASS" will be displayed.
	4 - If the exams average is less than 50, will ask that to student "Do you want to take the 3rd exam?".
	5 - If answer of the question is "YES" The average of 3 exams will be calculated,
		if this average is less than 50, the message will be "FAIL".
		If this average is more than 50, the message will be "PASS".
	6 - If answer of the question is "NO" the message will be "FAIL".

*/

func main() {

	var exam1, exam2, exam3, avg float64
	var answer string

	fmt.Print("Result of the first exam : ")
	fmt.Scanln(&exam1)
	fmt.Print("Result of the second exam : ")
	fmt.Scanln(&exam2)

	avg = (exam1 + exam2) / 2
	fmt.Printf("Avarage of 2 Exams : %v\n", avg)

	if avg >= 50 {
		fmt.Println("Status : PASSED")
	} else {
		fmt.Println("Status : FAILED")
	CONFIRMATION:
		fmt.Println("Do you want to take the 3rd exam? (YES : 'y' NO : 'n')")
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			fmt.Print("Result of the third exam : ")
			fmt.Scanln(&exam3)

			avg = (exam1 + exam2 + exam3) / 3
			fmt.Printf("Avarage of 3 Exams : %v\n", avg)

			if avg >= 50 {
				fmt.Println("Status : PASSED")
			} else {
				fmt.Println("Status : FAILED")
			}
		} else if answer == "n" || answer == "N" {
			fmt.Println("Status : FAILED")
		} else {
			fmt.Println("You Made the Wrong Choice.")
			goto CONFIRMATION
		}
	}
}
