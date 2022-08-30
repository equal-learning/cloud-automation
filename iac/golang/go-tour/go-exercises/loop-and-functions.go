package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	z := float64(1) // first guess

	for i := 0; i < 10; i++ {

		z -= (z*z - x) / (2 * z) // continuous improvement
		fmt.Println(z)

	}

	return z

}

func main() {
	fmt.Println(Sqrt(16))
}
