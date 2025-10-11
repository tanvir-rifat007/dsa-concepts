// Sum all numbers from 0 to n-1 that are NOT divisible by 3 or 5.
// For n=15: We want 1+2+4+7+8+11+13+14 = 60

package main

import "fmt"

// naive approach
func calculateFizzbuzzNumberNaive(n int) int {
	sum := 0

	for i := 0; i < n; i++ {
		if i%3 == 0 {
			continue
		}
		if i%5 == 0 {
			continue
		}
		if i%3 == 0 && i%5 == 0 {
			continue
		}

		sum += i

	}

	return sum

}

func calculateFizzbuzzNumber(n int) int {

	// sum of 0 to n-1 numbers
	// here n is (n-1) and the formula is : n * (n+1) / 2
	totalSum := (n * (n - 1)) / 2

	m3 := (n - 1) / 3
	m5 := (n - 1) / 5
	m15 := (n - 1) / 15 // this 15 divisible both 3 and 5

	S3 := 3 * (m3 * (m3 + 1) / 2)
	S5 := 5 * (m5 * (m5 + 1) / 2)
	Sn := 15 * (m15 * (m15 + 1) / 2)

	return totalSum - S3 - S5 + Sn

}

func main() {

	fmt.Printf("%d\n", calculateFizzbuzzNumber(10000000))

}
