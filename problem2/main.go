package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	alph := make(map[rune]int)

	alphString := "abcdefg界"

	for _, rune := range alphString {
		alph[rune]++
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your strings")

	for scanner.Scan() {
		result := checkString(alph, scanner.Text())
		fmt.Printf("Comparation result: %v\n", result)
	}
}

// We are passing each string rune with appropriate decrementing alphabet map value
// If any map value becomes negative - that seems input string is not properly fulfilled with provided alphabet
// Complexity of current approach is N + M where N is equal to length of alphabet, M is equal to length of input string
func checkString(alph map[rune]int, input string) bool {
	for _, rune := range input {
		alph[rune]--
		if alph[rune] < 0 {
			return false
		}
	}

	return true
}
