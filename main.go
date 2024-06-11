package main

import (
	"fmt"
	"strings"

	"github.com/saurabhkale77/calculator/addition"
	"github.com/saurabhkale77/calculator/division"
	"github.com/saurabhkale77/calculator/multiplication"
	"github.com/saurabhkale77/calculator/subtraction"
)

func main() {
	var a, b int

	defer func() {
		if r := recover(); r != nil {
			if strings.Contains(fmt.Sprintf("%v", r), "divide by zero") {
				fmt.Println("\nCan't divide by zero")
			} else {
				fmt.Println("\nUnknown issue")
			}
		} else {
			fmt.Println("\nNo issues")
		}
	}()

	fmt.Println("\nPlease enter first number")
	fmt.Scanln(&a)

	fmt.Println("\nPlease enter second number")
	fmt.Scanln(&b)

	fmt.Printf("\nAddition result:")
	fmt.Printf("\nAddition of %d and %d is %d\n", a, b, addition.Add(a, b))

	fmt.Printf("\nSubtraction result:")
	fmt.Printf("\nSubtraction of %d and %d is %d\n", a, b, subtraction.Subtract(a, b))

	fmt.Printf("\nMultiplication result:")
	fmt.Printf("\nMultiplication of %d and %d is %d\n", a, b, multiplication.Multiply(a, b))

	fmt.Printf("\nDivision result:")
	fmt.Printf("\nDivision of %d and %d is %.d\n", a, b, division.Divide(a, b))
}
