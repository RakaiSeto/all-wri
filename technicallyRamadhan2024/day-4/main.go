package main

func main() {

}

func ValidBraces(str string) bool {
	// given a string of braces, return true if the braces are valid, and false if they are not
	// create a stack to keep track of the braces
	stack := []rune{}
	// iterate through the string
	for _, v := range str {
		// if the brace is an opening brace, push it onto the stack
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			// if the brace is a closing brace, pop the last brace from the stack
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// if the popped brace does not match the closing brace, return false
			if (v == ')' && last != '(') || (v == ']' && last != '[') || (v == '}' && last != '{') {
				return false
			}
		}
	}
	// if the stack is not empty, return false
	if len(stack) != 0 {
		return false
	}
	// if the stack is empty, return true
	return true
}
