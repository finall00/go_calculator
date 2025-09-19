package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	Version = "1.0.0"
	PI      = 3.14159
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	programName := filepath.Base(os.Args[0])
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <add|sub|mult|div> <num1> <num2>\n", programName)
		return
	}

	num1Str := os.Args[2]
	num2Str := os.Args[3]

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Error: both arguments must be numbers.")
		return
	}

	var result float64
	var err error

	switch strings.ToLower(os.Args[1]) {
	case "add":
		result = float64(Add(int(num1), int(num2)))
	case "sub":
		result = float64(Sub(int(num1), int(num2)))
	case "mult":
		result = float64(Mul(int(num1), int(num2)))
	case "div":
		result, err = Div(num1, num2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	default:
		fmt.Println("Error: unknown operation. Use add, sub, mul, or div.")
		return
	}

	fmt.Printf("Result: %v\n", result)
}
