package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}

func main() {

	for {

		fmt.Print("Введите выражение (например, '2 + 2' или 'III * IV'): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Ошибка: Неверный формат ввода. Попробуйте снова.")
			continue
		}

		operand1 := parts[0]
		operator := parts[1]
		operand2 := parts[2]

		isArabic1 := isArabicNumber(operand1)
		isArabic2 := isArabicNumber(operand2)

		result := calculate(operand1, operator, operand2, isArabic1, isArabic2)

		fmt.Println("Результат:", result)
	}
}

// Проверка, арабское
func isArabicNumber(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func calculate(operand1, operator, operand2 string, isArabic1, isArabic2 bool) interface{} {
	num1, err1 := strconv.Atoi(operand1)
	num2, err2 := strconv.Atoi(operand2)

	if err1 == nil && err2 == nil {
		// Оба операнда арабские числа
		return calculateArabic(num1, operator, num2)
	} else if isArabic1 || isArabic2 {
		// Несовпадение типов ввода
		fmt.Println("Ошибка: Несоответствие типов ввода. Поддерживаются только арабские и римские числа (1-10). Попробуйте снова.")
		return nil
	} else {
		// Оба операнда римские числа
		return calculateRoman(operand1, operator, operand2)
	}
}

// вычисления арабские
func calculateArabic(num1 int, operator string, num2 int) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: Деление на ноль. Попробуйте снова.")
			return 0
		}
		return num1 / num2
	default:
		fmt.Println("Ошибка: Неизвестная операция. Поддерживаются только +, -, *, /")
		return 0
	}
}

// вычисления риские
func calculateRoman(operand1, operator, operand2 string) string {
	num1 := romanToArabic(operand1)
	num2 := romanToArabic(operand2)

	resultArabic := calculateArabic(num1, operator, num2)
	return arabicToRoman(resultArabic)
}

// преобразование нотации с рим в араб
func romanToArabic(romanStr string) int {
	result := 0
	prevValue := 0

	for i := len(romanStr) - 1; i >= 0; i-- {
		value := romanNumerals[rune(romanStr[i])]

		if value < prevValue {
			result -= value
		} else {
			result += value
		}

		prevValue = value
	}

	return result
}

// обратное преобр
func arabicToRoman(num int) string {
	if num <= 0 || num > 3999 {
		return "Недопустимое значение для римского числа"
	}

	var result strings.Builder

	// соответствие значения числа его римскому представлению
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
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

	for _, numeral := range romanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}
