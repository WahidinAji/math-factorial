package dimension

import "fmt"

func Dimension() {
	var larik [2]int
	var larik2 [10]float32

	larik[0] = 10
	larik[1] = 20
	larik2[3] = 4444.4
	larik2[5] = 6666.6
	fmt.Print(larik[0],larik[1],"\n")
	fmt.Print(larik2[3],larik2[5])
}
