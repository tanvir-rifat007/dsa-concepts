package main

import (
	"fmt"
	"strings"
)

func cleanPath(p string) string {
	abs := strings.HasPrefix(p, "/")
	parts := []string{}

	for _, part := range strings.Split(p, "/") {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if len(parts) > 0 {
				parts = parts[:len(parts)-1]
			}
			continue
		}
		parts = append(parts, part)
	}

	res := strings.Join(parts, "/") //  etc/bar/baz.txt
	if abs {
		res = "/" + res //  /etc/bar/baz.txt
	}
	return res
}

func main() {
	fmt.Println(cleanPath("/etc/foo/../bar/baz.txt"))
}
