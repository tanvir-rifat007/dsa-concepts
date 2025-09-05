package main

import (
	"fmt"
	"strconv"
	"strings"
)

func smallestNumber(n, t int) int {
	for {

		parts := strings.Split(strconv.Itoa(n), "")

		digitsProduct := 1
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			digitsProduct *= num
		}

		if digitsProduct%t == 0 {
			break

		}
		digitsProduct = 1
		n++

	}
	return n
}

func main() {
	res := smallestNumber(15, 3)
	fmt.Println(res)

}
