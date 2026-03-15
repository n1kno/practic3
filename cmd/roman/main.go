package main

import (
	"fmt"
	"strings"
)

func romanToArabic(roman string) int {
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	roman = strings.ToUpper(roman)
	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currentValue := values[rune(roman[i])]
		
		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}
		prevValue = currentValue
	}

	return result
}

func main() {
	examples := []string{"III", "IV", "IX", "LVIII", "MCMXC", "MMXXIV"}
	
	for _, roman := range examples {
		arabic := romanToArabic(roman)
		fmt.Printf("%s = %d\n", roman, arabic)
	}
}
