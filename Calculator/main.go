package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculate struct {
	a      int
	b      int
	result func(i int, j int) int
}

func main() {
	fmt.Println("Calculator")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter first value")
	a, _ := reader.ReadString('\n')
	input1, _ := strconv.Atoi(strings.TrimSpace(a))

	fmt.Println("Enter second value")
	b, _ := reader.ReadString('\n')
	input2, _ := strconv.Atoi(strings.TrimSpace(b))

	fmt.Println("Enter choice \n1. Addition \n2. Substraction \n3. Multiplication\n4.Division")
	c, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(c))

	// calculateSimple(input1, input2, choice)
	calculateWithStruct(input1, input2, choice)
}

func calculateSimple(a int, b int, intChoice int) {
	switch intChoice {
	case 1:
		fmt.Println("Result of Addition: ", a+b)
	case 2:
		if a > b {
			fmt.Println("Result of Substraction: ", a-b)
		} else {
			fmt.Println("Result of Substraction: ", b-a)
		}
	case 3:
		fmt.Println("Result of Multiplication: ", a*b)
	case 4:
		if b == 0 {
			fmt.Println("Wrong operational values")
		} else {
			fmt.Println("Result of Division: ", float32(a)/float32(b))
		}
	}
}

func calculateWithStruct(x int, y int, choice int) {
	switch choice {
	case 1:
		answer := Calculate{x, y, func(c int, d int) int {
			return c + d
		}}
		fmt.Println("Result of Addition: ", answer.result(answer.a, answer.b))
	case 2:
		if x > y {
			answer := Calculate{x, y, func(c int, d int) int {
				return c - d
			}}
			fmt.Println("Result of Substraction: ", answer.result(answer.a, answer.b))
		} else {
			answer := Calculate{x, y, func(c int, d int) int {
				return d - c
			}}
			fmt.Println("Result of Substraction: ", answer.result(answer.a, answer.b))
		}
	case 3:
		answer := Calculate{x, y, func(c int, d int) int {
			return c * d
		}}
		fmt.Println("Result of Multiplication: ", answer.result(answer.a, answer.b))
	case 4:
		if y == 0 {
			fmt.Println("Wrong operational values")
		} else {
			answer := Calculate{x, y, func(c int, d int) int {
				return c / d
			}}
			fmt.Println("Result of Division: ", answer.result(answer.a, answer.b))
		}
	}
}
