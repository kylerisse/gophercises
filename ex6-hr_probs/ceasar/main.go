package main

import "fmt"

func main() {
	var text string
	var length, offset int
	_, err := fmt.Scanf("%d\n%s\n%d\n", &length, &text, &offset)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(shift(text, offset))
}

func shift(word string, offset int) string {
	newWord := []byte{}
	for _, c := range word {
		newWord = append(newWord, byte(charShift(rune(c), offset)))
	}
	return string(newWord)
}

func charShift(char rune, offset int) rune {
	// uppercase
	if char >= 65 && char <= 90 {
		return rune((((int(char) + offset) - 65) % 26) + 65)
	}
	// lowercase
	if char >= 97 && char <= 122 {
		return rune((((int(char) + offset) - 97) % 26) + 97)
	}
	// not letter
	return char
}
