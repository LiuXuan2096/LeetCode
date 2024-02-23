package Go_LeetCode

import "fmt"

func intToRoman(num int) string {
	romanChars := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}

	roman := fmt.Sprintf("%s%s%s%s",
		romanChars[3][num/1000%10],
		romanChars[2][num/100%10],
		romanChars[1][num/10%10],
		romanChars[0][num%10])

	return roman
}
