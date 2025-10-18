package main

import (
	"fmt"
)

func calculate(s string) int {
	stack := []int{}
	num := 0
	sign := '+'

	for i := 0; i < len(s); i++ {
		char := s[i]

		// If the character is a digit, build the current number
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		}

		// If the character is an operator or we're at the end of the string

		if (char != ' ' && char != '+' && char != '-' && char != '*' && char != '/' && char < '0') ||
			(char == '+' || char == '-' || char == '*' || char == '/') ||
			i == len(s)-1 {

			// Process the previous number with its sign
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}

			if char == '+' || char == '-' || char == '*' || char == '/' {
				sign = rune(char)
			}

			num = 0
		}
	}

	result := 0
	for _, val := range stack {
		result += val
	}

	return result
}

func main() {

	fmt.Printf("%d\n", calculate("42"))

}
