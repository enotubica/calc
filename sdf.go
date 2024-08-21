package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRoman = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

func romanToInt(roman string) (int, error) {
	if value, exists := romanNumerals[roman]; exists {
		return value, nil
	}
	return 0, errors.New("некорректные римские цифры")
}

func intToRomanNum(num int) (string, error) {
	if roman, exists := intToRoman[num]; exists {
		return roman, nil
	}
	return "", errors.New("результат выходит за пределы римских цифр")
}

func calculate(expression string) (string, error) {
	tokens := strings.Fields(expression)

	if len(tokens) != 3 {
		return "", errors.New("некорректное выражение. Введите выражение в форме 'a + b'.")
	}

	num1 := tokens[0]
	operator := tokens[1]
	num2 := tokens[2]

	// Проверка формата чисел
	isRoman := func() bool {
		_, exists1 := romanNumerals[num1]
		_, exists2 := romanNumerals[num2]
		return exists1 && exists2
	}()

	var value1, value2 int
	var err error

	if isRoman {
		value1, err = romanToInt(num1)
		if err != nil {
			return "", err
		}
		value2, err = romanToInt(num2)
		if err != nil {
			return "", err
		}
	} else {
		value1, err = strconv.Atoi(num1)
		if err != nil {
			return "", errors.New("некорректный формат арабских чисел")
		}
		value2, err = strconv.Atoi(num2)
		if err != nil {
			return "", errors.New("некорректный формат арабских чисел")
		}
	}

	if (value1 < 1 || value1 > 10) || (value2 < 1 || value2 > 10) {
		return "", errors.New("числа должны быть в диапазоне от 1 до 10")
	}

	var result int
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			return "", errors.New("деление на ноль")
		}
		result = value1 / value2
	default:
		return "", errors.New("некорректный оператор. Ожидается '+', '-', '*', '/'")
	}

	if isRoman {
		if result < 1 {
			return "", errors.New("результат не может быть представлен римскими цифрами")
		}
		return intToRomanNum(result)
	}
	return strconv.Itoa(result), nil
}

func main() {
	var expression string
	fmt.Println("Введите арифметическое выражение (например, 'VII + V' или '3 * 2'):")

	_, err := fmt.Scanln(&expression)
	if err != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}
