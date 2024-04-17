package piscine 

func TrimAtoi(s string) int {
	runes := []rune(s)
	sign := 1
	result := 0
	foundDigit := false
	for _, r := range runes {
		if unicode.IsDigit(r) {
			foundDigit = true
			result = result*10 + int(r-'0')
		} else if r == '-' && !foundDigit {
			sign = -1
		}
	}
	return result * sign
}
