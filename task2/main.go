package main

import (
	"fmt"
	"strings"
)

func romanToInt(roman string) (int, error) {
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	roman = strings.ToUpper(strings.TrimSpace(roman))

	if roman == "" {
		return 0, fmt.Errorf("пустая строка")
	}

	symbols := []rune(roman)
	result := 0

	for i, symbol := range symbols {
		value, exists := values[symbol]
		if !exists {
			return 0, fmt.Errorf("неизвестный символ: %c", symbol)
		}

		if i+1 < len(symbols) {
			nextValue, exists := values[symbols[i+1]]
			if !exists {
				return 0, fmt.Errorf(
					"неизвестный символ: %c", symbols[i+1],
				)
			}

			if value < nextValue {
				result -= value
				continue
			}
		}

		result += value
	}

	if result < 1 || result > 3999 {
		return 0, fmt.Errorf("число должно быть от 1 до 3999")
	}

	if intToRoman(result) != roman {
		return 0, fmt.Errorf("неправильная запись римского числа")
	}

	return result, nil
}

func intToRoman(number int) string {
	values := []int{
		1000, 900, 500, 400, 100, 90,
		50, 40, 10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD", "C", "XC",
		"L", "XL", "X", "IX", "V", "IV", "I",
	}

	var result strings.Builder

	for i, value := range values {
		for number >= value {
			result.WriteString(symbols[i])
			number -= value
		}
	}

	return result.String()
}

func main() {
	var roman string

	fmt.Print("Введите римское число: ")
	fmt.Scanln(&roman)

	result, err := romanToInt(roman)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Арабское число:", result)
}
