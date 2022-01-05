package main

import (
	"fmt"
	"math/factorial/dimension"
	"math/factorial/factorial"
)

func main() {
	dimension.Dimension()
	fmt.Print("\n kasih jarak factorial 1 \n")
	fmt.Println(factorial.Factorial(1))
	fmt.Print("\n kasih jarak factorial 5 \n")
	fmt.Println(factorial.Factorial(5))
	fmt.Println(factorial.FactorialTailRecursiveOddNumber(1,5))
	fmt.Println(factorial.FactorialTailRecursiveEvenNumber(1,4))
}

