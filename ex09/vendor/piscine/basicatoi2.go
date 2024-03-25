package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func BasicAtoi2(s string) int {
	runes := []rune(s)
	length := StrLen(s)
	result := 0

	for i := 0; i < length; i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			digit := int(runes[i] - '0')
			result = result*10 + digit
		} else {
			return 0
		}
	}
	return result

}

