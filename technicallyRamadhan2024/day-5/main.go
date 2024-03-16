package main

func main() {

}

func Decode(roman string) int {
	// given a roman numeral, return the value of the numeral
	// create a map of roman numerals to
	// iterate through the string
	// if the current numeral is smaller than the next numeral, subtract the current numeral from the total
	// if the current numeral is larger than the next numeral, add the current numeral to the total
	// return the total

	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	total := 0
	for i := 0; i < len(roman); i++ {
		if i < len(roman)-1 && romanMap[string(roman[i])] < romanMap[string(roman[i+1])] {
			total -= romanMap[string(roman[i])]
		} else {
			total += romanMap[string(roman[i])]
		}
	}
	return total
}
