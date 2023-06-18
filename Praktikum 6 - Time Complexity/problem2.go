package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}

func pow(x, n int) int {
	var result float64 = 0
	result = math.Pow(float64(x), float64(n))
	return int(result)
}
