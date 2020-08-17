package main

import (
	"strconv"
	"fmt"
)

func main() {
	for i := 1 ; i < 100 ; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(x int) string {
	if x % 15 == 0 {
		return "FizzBuzz"
	} else if x % 5 == 0 {
		return "Buzz"
	} else if x % 3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(x)
}