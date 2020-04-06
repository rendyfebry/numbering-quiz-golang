package main

import (
	"fmt"

	"github.com/rendyfebry/numbering-quiz-golang/lib/numbering"
)

func main() {
	num := numbering.NewNumber()

	fmt.Println("Sum(3,4):", num.Sum(3, 4))
	fmt.Println("Multiply(3,4):", num.Multiply(3, 4))
}
