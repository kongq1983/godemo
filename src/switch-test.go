package main

import "fmt"

func main() {

	printGrade("C")
	printGrade("F")
	printGrade2(88)

	fmt.Println()

	printGrade3("a")
	printGrade3("A")
	printGrade3("b")
	printGrade3("B")
	printGrade3("c")
	printGrade3("C")
	printGrade3("d")
	printGrade3("D")
	printGrade3("E")

}

func printGrade(grade string) {

	switch {
	case grade == "A":
		fmt.Printf("A1 \n")
	case grade == "B":
		fmt.Printf("B1 \n")
	case grade == "C":
		fmt.Printf("C1 \n")
	case grade == "D":
		fmt.Printf("D1 \n")
	default:
		fmt.Printf("Other. \n")
	}

}

func printGrade2(grade int) {

	switch {
	case grade >= 90:
		fmt.Printf("A1 \n")
	case grade >= 80:
		fmt.Printf("B1 \n")
	case grade >= 70:
		fmt.Printf("C1 \n")
	case grade >= 60:
		fmt.Printf("D1 \n")
	default:
		fmt.Printf("Other. \n")
	}

}

func printGrade3(grade string) {
	switch {
	case grade == "A", grade == "a":
		fmt.Printf("A1 grade=%s\n", grade)
	case grade == "b":
	case grade == "B":
		fmt.Printf("B1  grade=%s\n", grade)
	case grade == "c":
		fallthrough
	case grade == "C":
		fmt.Printf("C1 grade=%s\n", grade)
	case grade == "D":
		fmt.Printf("D1 grade=%s\n", grade)
	default:
		fmt.Printf("Other. grade=%s\n", grade)
	}
}
