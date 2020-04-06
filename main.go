package main

import (
	"fmt"

	"github.com/rendyfebry/numbering-quiz-golang/lib/numbering"
)

func main() {
	num := numbering.NewNumber()

	fmt.Println("Sum(3,4): \t\t", num.Sum(3, 4))
	fmt.Println("Multiply(3,4): \t\t", num.Multiply(3, 4))
	fmt.Println("GetPrimes(4): \t\t", num.GetPrimes(4))
	fmt.Println("GetFibonnacy(4): \t", num.GetFibonnacy(4))
}
