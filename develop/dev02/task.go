package main

import (
	"fmt"
	"strconv"
	 
	"unicode"
)

func unpackString(str string) (string, error) {
	data := []rune(str)
	result := make([]rune, 0)

	if len(data) == 0 {
		return "", nil  
	}

	if unicode.IsDigit(data[0]) {
		return "", fmt.Errorf("start with num")
	}

	i := 0
	for i < len(data) {
		if unicode.IsDigit(data[i]) {
			number := make([]rune, 0)
			for i < len(data) && unicode.IsDigit(data[i]) {
				number = append(number, data[i])
				i++
			}
			n, err := strconv.Atoi(string(number))
			if err != nil {
				return "", fmt.Errorf("incorect format")
			}

			if len(result) == 0 {
				return "", fmt.Errorf("start with num")
			}

			lastChar := result[len(result)-1]
			for j := 0; j < n-1; j++ {
				result = append(result, lastChar)
			}
		} else {
			result = append(result, data[i])
		}
		i++
	}

	return string(result), nil
}

func main() {
	inputs := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
		"qwe\\\\5",
	}

	for _, input := range inputs {
		result, err := unpackString(input)
		if err != nil {
			println("error:", err.Error())
		} else {
			println("input:", input)
			println("output:", result)
		}
	}
}
