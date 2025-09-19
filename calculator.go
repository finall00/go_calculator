package main

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

func Div(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
