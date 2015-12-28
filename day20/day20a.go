/*

1 * 10 = 10
1 * 10 + 2 * 10 = 30
1 * 10 + 3 * 10 = 40
1 * 10 + 2 * 10 + 4 * 10 = 70
1 * 10

*/
package main

import (
	"fmt"
	"math"
)

func compute(nr int) bool {
	sum := 0
	for i := 1; i <= int(math.Sqrt(float64(nr))); i += 1 {
		if nr%i == 0 {
			sum += i
			sum += int(nr / i)
		}
	}
	if sum >= 3310000 {
		fmt.Println(sum, nr)
		return false
	}
	return true
}

func main() {
	// start from 100k, just to reduce the space
	for i := 100000; compute(i); i += 1 {
	}
}
