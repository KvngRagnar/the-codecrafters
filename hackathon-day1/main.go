package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) (int, string) {
	if b == 0 {
		return 0, "denominator cannot be 0"
	}
	return a / b, " "
}

func colour(text, color string) string {
	return color + text + "\033[33m"
}

func main() {
	for {
		var a, b int
		fmt.Println(colour("input first integer", "\033[33m"))
		fmt.Scanln(&a)

		fmt.Println("input second integer", "\033[33m")
		fmt.Scanln(&b)

		var operation string

		fmt.Println("choose an operation (add, sub, mul, div, quit, help)")
		fmt.Scanln(&operation)

		switch operation {
		case "add":
			fmt.Println(add(a, b))
			continue

		case "sub":
			fmt.Println(sub(a, b))
			continue

		case "mul":
			fmt.Println(mul(a, b))
			continue

		case "div":
			fmt.Println(div(a, b))
			continue

		case "quit":
			fmt.Println("Thank you for testing my code, abientot!!")
			return

		case "help":
			fmt.Println("supported commands = add, sub, mul, div, quit, help")

		default:
			fmt.Println("Invalid expression, type help for valid expressions", "\033[31m")
		}

	}

}
