package main

import (
	"fmt"
	"strings"
)

func normalizeURL(path string) string {
	parts := strings.Split(path, "/")

	var stack []string

	for _, part := range parts {
		if part == "" {

			if len(stack) == 0 {
				stack = append(stack, "")

			}
		} else if part == ".." {
			if len(stack) > 0 && stack[len(stack)-1] != "" {

				stack = stack[:len(stack)-1]

			}

		} else if part != "." {
			stack = append(stack, part)
		}

	}

	result := strings.Join(stack, "/")

	if result == "" {
		return "/"

	}

	return result

}

func main() {

	url := "/etc/foo/../bar/baz"

	res := normalizeURL(url)

	fmt.Printf("%s\n", res)
}
