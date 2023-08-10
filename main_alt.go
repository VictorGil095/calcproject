package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calc \"expression\"")
		return
	}

	left := os.Args[1]
	operator := os.Args[2]
	right := os.Args[3]

	var result string

	switch operator {
	case "+":
		result = addStrings(left, right)
	case "-":
		result = subtractStrings(left, right)
	case "*":
		if num, err := strconv.Atoi(right); err == nil {
			result = multiplyString(left, num)
		} else {
			fmt.Println("Invalid expression format.")
			return
		}
	case "/":
		if num, err := strconv.Atoi(right); err == nil {
			result = divideString(left, num)
		} else {
			fmt.Println("Invalid expression format.")
			return
		}
	default:
		fmt.Println("Unsupported operator.")
		return
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Println(result)
}

func addStrings(a, b string) string {
	return a + b
}

func subtractStrings(a, b string) string {
	return strings.Replace(a, b, "", 1)
}

func multiplyString(s string, n int) string {
	var result strings.Builder
	for i := 0; i < n; i++ {
		result.WriteString(s)
	}
	return result.String()
}

func divideString(s string, n int) string {
	return s[:len(s)/n]
}
