package main

func validParenthesis(s []string) bool {

	if len(s) == 0 {
		return false
	}

	var res []string

	for i, _ := range s {
		if s[i] == "(" {
			res = append(res, ")")

		} else if s[i] == "[" {

			res = append(res, "]")
		} else if s[i] == "{" {
			res = append(res, "}")

		} else {
			if len(res) == 0 || res[len(res)-1] != s[i] {
				return false
			}
			res = res[:len(res)-1]
		}

	}

	return len(res) == 0
}

func main() {

	s := []string{"(", "[", "]", ")"}
	validParenthesis(s)

}
