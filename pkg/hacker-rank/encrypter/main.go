package main

import "fmt"

func removeNumbers(s string) string {
	var result string
	for _, c := range s {
		if c < '0' || c > '9' {
			result += string(c)
		}
	}
	return result
}

func trim(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			result = s[i:]
			break
		}
	}
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != ' ' {
			result = result[:i+1]
			break
		}
	}
	return result
}

func reverse(s string) string {
	var result string
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func encrypt(s string) string {
	var result string
	result = trim(s)
	result = removeNumbers(result)
	result = reverse(result)

	return result
}

func main() {
	fmt.Println(encrypt("oll123eH56"))
}
