package main

import "fmt"

func gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}

	return a

}

func canMeasureWater(x int, y int, target int) bool {

	return x+y >= target && target%gcd(x, y) == 0

}

func main() {

	fmt.Printf("%v\n", canMeasureWater(3, 5, 4))

}
