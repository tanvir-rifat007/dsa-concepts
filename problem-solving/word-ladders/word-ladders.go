package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {

	wordSet := make(map[string]bool)

	for _, word := range wordList {
		wordSet[word] = true

	}

	if !wordSet[endWord] {
		return 0
	}

	queue := []string{beginWord}
	level := 1

	for len(queue) > 0 {

		size := len(queue)

		for i := 0; i < size; i++ {
			currentWord := queue[0]
			queue = queue[1:]

			if currentWord == endWord {

				return level
			}

			wordBytes := []byte(currentWord)

			for j := 0; j < len(wordBytes); j++ {
				originalChar := wordBytes[j]
				fmt.Println("char in byte: ,char in string", originalChar, string(originalChar))
				for ch := 'a'; ch <= 'z'; ch++ {
					wordBytes[j] = byte(ch)
					newWord := string(wordBytes)
					fmt.Println(newWord)

					if wordSet[newWord] {
						queue = append(queue, newWord)
						delete(wordSet, newWord)
					}
				}

				// Restore original character
				wordBytes[j] = originalChar

			}

		}
		level++

	}

	return 0

}

func main() {

	beginWord1 := "hit"
	endWord1 := "cog"
	wordList1 := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	result1 := ladderLength(beginWord1, endWord1, wordList1)
	fmt.Printf("Example 1: %d\n", result1) // Output: 5

}
