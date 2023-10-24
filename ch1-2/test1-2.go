package main

import (
	"fmt"
	"os"
)

func main() {
	for key, value := range os.Args {
		fmt.Printf("%v:%v\n", key, value)
	}
}
