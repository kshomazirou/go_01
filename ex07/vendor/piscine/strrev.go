package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func StrRev(s string) string {
	runes := []rune(s)
	length := StrLen(s)

	for i,j := 0, length-1; i < j; i, j =i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversedString := string(runes)
	return reversedString
}
