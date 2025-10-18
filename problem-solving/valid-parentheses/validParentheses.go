package main

import "fmt"

func validParentheses(s string) bool {

	if len(s) == 0 {
		return false
	}
	var res []rune

	for _, ch := range s {

		if ch == '(' {
			res = append(res, ')')
		} else if ch == '{' {
			res = append(res, '}')
		} else if ch == '[' {
			res = append(res, ']')
		} else {
			if len(res) == 0 || res[len(res)-1] != ch {

				return false
			}

			res = res[:len(res)-1]

		}

	}

	return len(res) == 0

}

func main() {

	fmt.Println(validParentheses("((()])"))
}
