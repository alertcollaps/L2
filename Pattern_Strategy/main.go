package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Subtract(a, b int) int {
	return a - b
}

type Strategy struct {
}

func (st Strategy) execute(f func(int, int) int, a int, b int) int {
	return f(a, b)
}

func main() {
	st := new(Strategy)
	result := st.execute(Multiply, 4, 5)
	fmt.Println(result)
}
