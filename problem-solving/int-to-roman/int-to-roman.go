package main

import "fmt"

type RomanRepresentation struct {
	value int
	roman string
}

func intToRoman(n int) string {

	romans := []RomanRepresentation{

		{value: 1000, roman: "M"},

		{value: 900, roman: "CM"}, // this means M(1000) - C(100) = 900(CM)
		{value: 500, roman: "D"},
		{value: 400, roman: "CD"},

		{value: 100, roman: "C"},

		{value: 90, roman: "XC"},

		{value: 50, roman: "L"},

		{value: 40, roman: "XL"},

		{value: 10, roman: "X"},

		{value: 9, roman: "IX"},

		{value: 5, roman: "V"},

		{value: 4, roman: "IV"},

		{value: 1, roman: "I"},
	}

	for _, elm := range romans {
		if elm.value <= n {

			return elm.roman + intToRoman(n-elm.value)

		}

	}

	return ""

}

func main() {
	fmt.Printf("The result in roman from int : %s\n", intToRoman(3375))

}
