package main

import (
	"C"
	"fmt"
	"os"
	"strconv"
)

/** NOTE: the //export sum line is necessary here! **/
//export sum
func sum(number1, number2 int) int {
	return variadicSum(number1, number2)
}

func variadicSum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	arguments := os.Args[1:]
	sliceLength := len(arguments)
	integers := make([]int, sliceLength)

	for i, argument := range arguments {
		integer, _ := strconv.Atoi(argument) // ignore err in example
		integers[i] = integer
	}

	fmt.Println(variadicSum(integers...))
}
