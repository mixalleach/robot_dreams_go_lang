package main

import (
	"fmt"
	"strconv"
)

func FibonacciIterative(n int) int {
	// Функція вираховує і повертає n-не число фібоначчі
	// Імплементація без використання рекурсії

	if n < 2 {
		return n
	}

	first := 0
	second := 1
	fibonacci := 1

	for i := 2; i <= n; i++ {
		fibonacci = first + second
		first = second
		second = fibonacci
	}

	return fibonacci
}

func FibonacciRecursive(n int) int {
	// Функція вираховує і повертає n-не число фібоначчі
	// Імплементація з використанням рекурсії

	if n < 2 {
		return n
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func IsPrime(n int) bool {
	// Функція повертає `true` якщо число `n` - просте.
	// Інакше функція повертає `false`

	if n <= 1 {
		return true
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsBinaryPalindrome(n int) bool {
	// Функція повертає `true` якщо число `n` у бінарному вигляді є паліндромом
	// Інакше функція повертає `false`
	//
	// Приклади:
	// Число 7 (111) - паліндром, повертаємо `true`
	// Число 5 (101) - паліндром, повертаємо `true`
	// Число 6 (110) - не є паліндромом, повертаємо `false`

	binary := strconv.FormatInt(int64(n), 2)
	length := len(binary)

	for i := 0; i < length/2; i++ {
		if binary[i] != binary[length-1-i] {
			return false
		}
	}

	return true
}

func ValidParentheses(s string) bool {
	// Функція повертає `true` якщо у вхідній стрічці дотримані усі правила високристання дужок
	// Правила:
	// 1. Допустимі дужки `(`, `[`, `{`, `)`, `]`, `}`
	// 2. У кожної відкритої дужки є відповідна закриваюча дужка того ж типу
	// 3. Закриваючі дужки стоять у правильному порядку
	//    "[{}]" - правильно
	//    "[{]}" - не правильно
	// 4. Кожна закриваюча дужка має відповідну відкриваючу дужку

	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}

			stack = stack[:len(stack)-1]
		default:
			continue
		}
	}

	return len(stack) == 0
}

func Increment(num string) int {
	// Функція на вхід отримує стрічку яка складається лише з символів `0` та `1`
	// Тобто стрічка містить певне число у бінарному вигляді
	// Потрібно повернути число на один більше

	n, _ := strconv.ParseInt(num, 2, 64)

	return int(n) + 1
}

func main() {
	//n := 10
	//fmt.Printf("FibonacciIterative from %d is %d \n", n, FibonacciIterative(n))
	//fmt.Printf("FibonacciRecursive from %d is %d \n", n, FibonacciRecursive(n))
	//fmt.Printf("Number %d is prime %t", n, IsPrime(n))
	//fmt.Printf("%d is binary palindrome: %t\n", n, IsBinaryPalindrome(n))

	//n := "00001010"
	//fmt.Printf("Increment of %s is %d\n", n, Increment(n))

	fmt.Printf("%t\n", ValidParentheses("[{]}"))
}
