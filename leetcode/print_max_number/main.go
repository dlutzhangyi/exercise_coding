package main

import (
	"fmt"
	"strconv"
)

func PrintToMaxOfNDigits(n int) {
	if n <= 0 {
		return
	}
	number := make([]string, n)
	for i := 0; i < 10; i++ {
		number[0] = strconv.Itoa(i + 0)
		PrintToMaxOfNDigitsRecursive(number, n, 0)
	}
}

func PrintToMaxOfNDigitsRecursive(number []string, length int, index int) {
	if index == length-1 {
		fmt.Println(number)
		return
	}
	for i := 0; i < 10; i++ {
		number[index+1] = strconv.Itoa(i + 0)
		PrintToMaxOfNDigitsRecursive(number, length, index+1)
	}
}

func main() {
	PrintToMaxOfNDigits(2)
	//PrintToMaxOfNDigits(2)
	//PrintToMaxOfNDigits(3)
}
