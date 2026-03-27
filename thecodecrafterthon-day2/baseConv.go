package main

import (
	"fmt"
	"strconv"
)

func ConvToDec(num string, base int) (int64, error) {
	decimal, err := strconv.ParseInt(num, base, 64)
	return decimal, err
}

func ConvToHex(num string) (int64, error) {
	hex, err := strconv.ParseInt(num, 16, 64)
	return hex, err
}

func ConvToBin(num string) (int64, error) {
	bin, err := strconv.ParseInt(num, 2, 64)
	return bin, err
}

func main() {
	for {
		var num string
		fmt.Println("Input string to convet")
		fmt.Scanln(&num)

		var operation string
		fmt.Println("choose operation (bin, dec, hex)")
		fmt.Scanln(&operation)

		switch operation {
		case "bin":
			fmt.Println(ConvToDec(num, 64))
			continue
		}
	}
}
