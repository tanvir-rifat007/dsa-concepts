// o(n*2)

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type RomanRepresentation struct {
// 	value int
// 	roman string
// }

// func romanToInt(str string) int {
// 	romans := []RomanRepresentation{
// 		{1000, "M"},
// 		{900, "CM"},
// 		{500, "D"},
// 		{400, "CD"},
// 		{100, "C"},
// 		{90, "XC"},
// 		{50, "L"},
// 		{40, "XL"},
// 		{10, "X"},
// 		{9, "IX"},
// 		{5, "V"},
// 		{4, "IV"},
// 		{1, "I"},
// 	}

// 	res := 0
// 	for _, rn := range romans {
// 		for strings.HasPrefix(str, rn.roman) {
// 			res += rn.value
// 			str = str[len(rn.roman):] // remove the matched prefix
// 		}
// 	}

// 	return res
// }

// func main() {
// 	fmt.Printf("The result of MCMXCIV in integer is: %d\n", romanToInt("MCMXCIV"))
// }

// O(n)
package main

import "fmt"

func romanToInt(str string) int {

	romans := map[string]int{

		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var res int

	for i := 0; i < len(str); i++ {
		if i < len(str)-1 && romans[string(str[i])] < romans[string(str[i+1])] {

			res += romans[string(str[i+1])] - romans[string(str[i])]
			i++

		} else {
			res += romans[string(str[i])]
		}

	}

	return res

}

func main() {

	fmt.Printf("Roman to int: %d\n", romanToInt("XLV"))
}
