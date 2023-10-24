package main

import "fmt"

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	fmt.Printf("the %v th fib number is %v \n", 10, fib(10))
	fmt.Printf("the gcd of %v and %v is %v \n", 20, 17, gcd(20, 17))
}
