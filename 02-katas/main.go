package main

import (
	"fmt"

	"github.com/course-go/exercises/02-katas/internal/katas"
)

func main() {
	fmt.Println(katas.Fizzbuzz(0))  // "0"
	fmt.Println(katas.Fizzbuzz(3))  // "fizz"
	fmt.Println(katas.Fizzbuzz(5))  // "buzz"
	fmt.Println(katas.Fizzbuzz(15)) // "fizzbuzz"

	// You can test the code here if you'd like
}
