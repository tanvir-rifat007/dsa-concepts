package main

import "fmt"

/*

  for word 'anagram'
hash = {a:3,g:1,n:1,m:1};

for word 'nagaram'
hash = {a:3, g:1, n:1,m:1}

*/

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hash1 := make(map[byte]int)

	hash2 := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		hash1[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		hash2[t[i]]++
	}

	for char, count := range hash1 {

		if hash2[char] != count {

			return false
		}

	}
	return true

}

func main() {
	res := isAnagram("cat", "rat")
	fmt.Printf("%v\n", res)
}
