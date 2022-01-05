package factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 1,Factorial(0),"Must be 1")
	assert.Equal(t, 1,Factorial(1),"Must be 1")
	assert.Equal(t, 2,Factorial(2),"Must be 2")
	assert.Equal(t, 6,Factorial(3),"Must be 6")
	assert.Equal(t, 24,Factorial(4),"Must be 24")
	assert.Equal(t, 120,Factorial(5),"Must be 120")
}
func TestFactorialWithLooping(t *testing.T) {
	assert.Equal(t, 1,FactorialWithLooping(0),"Must be 1")
	assert.Equal(t, 1,FactorialWithLooping(1),"Must be 1")
	assert.Equal(t, 2,FactorialWithLooping(2),"Must be 2")
	assert.Equal(t, 6,FactorialWithLooping(3),"Must be 6")
	assert.Equal(t, 24,FactorialWithLooping(4),"Must be 24")
	assert.Equal(t, 120,FactorialWithLooping(5),"Must be 120")
}

func TestFactorialTailRecursive(t *testing.T) {
	assert.Equal(t, 1,FactorialTailRecursive(1,0),"Must be 1")
	assert.Equal(t, 1,FactorialTailRecursive(1,1),"Must be 1")
	assert.Equal(t, 2,FactorialTailRecursive(1,2),"Must be 2")
	assert.Equal(t, 6,FactorialTailRecursive(1,3),"Must be 6")
	assert.Equal(t, 24,FactorialTailRecursive(1,4),"Must be 24")
	assert.Equal(t, 120,FactorialTailRecursive(1,5),"Must be 120")
}
func TestFactorialTailRecursiveEvenNumber(t *testing.T) {
	//assert.Equal(t, 1,FactorialTailRecursiveEvenNumber(1,0),"Must be 1")
	//assert.Equal(t, 1,FactorialTailRecursiveEvenNumber(1,1),"Must be 1")
	assert.Equal(t, 2,FactorialTailRecursiveEvenNumber(1,2),"Must be 2")
	//assert.Equal(t, 6,FactorialTailRecursiveEvenNumber(1,3),"Must be 6")
	assert.Equal(t, 24,FactorialTailRecursiveEvenNumber(1,4),"Must be 24")
	//assert.Equal(t, 120,FactorialTailRecursiveEvenNumber(1,5),"Must be 120")
}
