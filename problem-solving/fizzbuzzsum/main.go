package main

import "fmt"

func calculateFizzbuzzNumber(n int) int {

	m3 := 15 / 3
	m5 := 15 / 5
	m15 := 15 / 15 // this 15 divisible both 3 and 5

	S3 := 3 * (m3 * (m3 + 1) / 2)
	S5 := 5 * (m5 * (m5 + 1) / 2)
	Sn := 15 * (m15 * (m15 + 1) / 2)

	return S3 + S5 - Sn

}

func main() {

	fmt.Printf("%d\n", calculateFizzbuzzNumber(15))

}
