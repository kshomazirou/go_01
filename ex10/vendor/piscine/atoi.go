package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func Atoi(s string) int {
	runes := []rune(s)
	length := StrLen(s)
	result := 0
	sign := 1
	flag := 0
	i := 0
	for j:=0;j < length; j++ {
		if runes[j] == '-' {
			sign *= -1
			flag++
		} else if runes[j] == '+' {
			flag++
		} else if runes[j] >= '0' && runes[j] <= '9' {
			break ;
		}
	}
	if flag == 1 {
		i = 1
	} else if flag > 1 {
		return 0
	}
	for ;i < length; i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			digit := int(runes[i] - '0')
			result = result*10 + digit
		} else {
			return 0
		}
	}
	return result * sign
}

