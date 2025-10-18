// problem: https://cses.fi/problemset/task/1068

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func weirdAlgorithm(n int) string {

	res := []string{strconv.Itoa(n)}

	for n != 1 {

		if n%2 == 0 {
			n = n / 2
			res = append(res, strconv.Itoa(n))
		} else if n%2 == 1 {
			n = (n * 3) + 1
			res = append(res, strconv.Itoa(n))
		} else {
			continue
		}

	}

	return strings.Join(res, " ")

}

func main() {

	fmt.Printf("%v\n", weirdAlgorithm(3))

}
