package main

import "strings"

var (
	intToRomanMap = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	num = 1994
)

func intToRoman(num int) string {
	romanStr := new(strings.Builder)
	for num > 0 {
		for _, pair := range intToRomanMap {
			if num >= pair.value {
				romanStr.WriteString(pair.symbol)
				num -= pair.value
				break
			}
		}
	}
	return romanStr.String()
}
