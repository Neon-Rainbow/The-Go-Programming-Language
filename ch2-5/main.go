package main

import (
	"ch2-5/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("%v Celsius is %v Fahrenheit \n", 100, tempconv.CToF(100))
	fmt.Printf("%v Fahrenheit is %v Celsius \n", 100, tempconv.FToC(100))
}
