package main

var (
	s        = "III"
	romanMap = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'C': 100,
		'L': 50,
		'D': 500,
		'M': 1000,
	}
)

func romanToInt(s string) int {
	total := 0
	prevValue := 0
	for i := 0; i < len(s); i++ {
		value := romanMap[s[i]]
		if prevValue < value {
			total += value - 2*prevValue
		} else {
			total += value
		}
		prevValue = value
	}
	return total
}
