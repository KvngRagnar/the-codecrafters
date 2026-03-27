package main

import (
	"fmt"
	"strconv"
)

func ConvToDec(num string, base int) (int64, error) {
	decimal, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		fmt.Println("Value is not a decimal number, kindly input a decimal number")
	}
	return decimal, err
}

func ConvToHex(num string) (int64, error) {
	hex, err := strconv.ParseInt(num, 16, 64)
	if err != nil {
		fmt.Println("Value is not a hexadecimal number, kindly input a decimal number")
	}
	return hex, err
}

func ConvToBin(num string) (int64, error) {
	bin, err := strconv.ParseInt(num, 2, 64)
	if err != nil {
		fmt.Println("Value is not a binary number, kindly input a decimal number")
	}
	return bin, err
}

func colour(text, color string) string {
	return color + text + "\033[0m"
}

func main() {
	for {
		var num string
		fmt.Println(colour("Input string to convet", "\033[34m"))
		fmt.Scanln(&num)

		var operation string
		fmt.Println(colour("choose operation (bin, dec, hex, quit)", "\033[36m"))
		fmt.Scanln(&operation)

		switch operation {
		case "bin":
			fmt.Println(ConvToBin(num))
			continue

		case "dec":
			fmt.Println(ConvToDec(num, 64))
			continue

		case "hex":
			fmt.Println(ConvToHex(num))
			continue

		case "quit":
			fmt.Println("Leaving so soon?, hope to see you soon")
			return

		default:
			fmt.Println("Not a valid conversion, try again")
			fmt.Println("valid conversions are: bin, dec and hex")
		}
	}
}
