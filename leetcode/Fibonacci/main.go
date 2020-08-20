package main

import "fmt"

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	pre1, pre2, res := 0, 1, 0
	for i := 2; i <= N; i++ {
		res = pre1 + pre2
		pre1 = pre2
		pre2 = res
	}
	return res
}

func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
}
