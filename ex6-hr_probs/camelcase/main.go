package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	_, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println(countWords(input))
}

func countWords(str string) int {
	count := 1
	for _, char := range str {
		if char == unicode.ToUpper(rune(char)) {
			count++
		}
	}
	return count
}
