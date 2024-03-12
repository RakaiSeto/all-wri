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

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
