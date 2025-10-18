package main

import (
	"fmt"
	"strconv"
)

func superPow(a int, b []int) int {
	var str string

	for i := 0; i < len(b); i++ {

		str += strconv.Itoa(b[i])
	}

	num, _ := strconv.Atoi(str)

	return num

}

func main() {

	fmt.Printf("%v %T\n", superPow(1, []int{4, 3, 3, 8, 5, 2}), superPow(1, []int{4, 3, 3, 8, 5, 2}))

}
