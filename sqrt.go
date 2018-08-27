package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	z := float64(1)
	zprev := z
	z -= (z*z - x) / (2 * z)

	for z != zprev {
		zprev = z
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z)
		// fmt.Println("XXX")
	}
	return zprev
}

func main() {
	fmt.Println(sqrt(144))
}
