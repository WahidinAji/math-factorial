package factorial

import "fmt"

func Factorial(value int) int {
	if value > 1 {
		total := value * Factorial(value-1)
		return total
	} else {
		return 1
	}
}

func FactorialWithLooping(value int) int {
	var result int = 1
	if value > 1 {
		for i := value; i >= 1; i-- {
			result *= i
		}
	}
	return result
}

func FactorialTailRecursive(total int, value int) int {
	if value > 1 {
		return FactorialTailRecursive(total * value, value-1);
	} else {
		return total
	}
}

func FactorialTailRecursiveEvenNumber(total int, value int) int {
	var result int = total
	if value > 1 {
		if value%2 ==0  {
			fmt.Println("====Even Number====")
			result = FactorialTailRecursive(total * value, value-1)
		}
		return result
	}
	return total
}
func FactorialTailRecursiveOddNumber(total int, value int) int {
	var result int = total
	if value > 1 {
		if value%2 != 0  {
			fmt.Println("====Odd Number====")
			result = FactorialTailRecursive(total * value, value-1)
		}
		return result
	}
	return total
}