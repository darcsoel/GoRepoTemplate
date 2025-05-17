package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	sum := 0

	for i := range 100_000_000 {
		sum += i
	}

	fmt.Println(sum)
}

func strConcat(first string, second string) string {
	return first + second
}
