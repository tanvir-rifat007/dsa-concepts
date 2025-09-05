package main

import "fmt"

func binpow(a float64, b int64) float64 {
	/*
		 recursive approach
			if b == 0 {
				return 1.0
			}
			if b < 0 {
				return 1.0 / binpow(a, -b)

			}

			res := binpow(a, b/2)

			if b%2 == 1 {
				return res * res * a
			}
			return res * res
	*/

	// this is the iterative approach

	// time complexity O(log b)
	// space complexity O(1)
	result := 1.0

	if b < 0 {
		return 1.0 / binpow(a, -b)

	}

	for b > 0 {
		if b%2 == 1 {
			result *= a

		}
		a *= a
		b >>= 1 // b = b / 2

	}

	return result

}

func main() {
	fmt.Println(binpow(3, 13))
	fmt.Println(binpow(2, 10))

	fmt.Println(binpow(2.0, -3))
}
