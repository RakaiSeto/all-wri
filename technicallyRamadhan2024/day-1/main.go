package kata

import "strings"

func SpinWords(str string) string {
	// given a sentence, reverse the words that have 5 or more letters, then return the reassembled sentence
	// split the sentence into words
	words := strings.Split(str, " ")
	// iterate through the words
	for i, word := range words {
		// if the word has 5 or more letters
		if len(word) >= 5 {
			// reverse the word
			words[i] = reverse(word)
		}
	}
	// reassemble the sentence
	return strings.Join(words, " ")
} // SpinWords

func reverse(s string) string {
	// given a string, reverse the string
	// convert the string to a slice of runes
	runes := []rune(s)
	// iterate through the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// swap the runes
		runes[i], runes[j] = runes[j], runes[i]
	}
	// convert the slice of runes back to a string
	return string(runes)
}
