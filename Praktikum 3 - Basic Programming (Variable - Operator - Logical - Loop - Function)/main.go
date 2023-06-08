package main

import (
	"fmt"
)

func main() {

	//	PROBLEM  1
	// var r float32
	// var t float32

	// fmt.Printf("Radius : ")
	// fmt.Scanf("%f\n", &r)
	// fmt.Printf("Tinggi : ")
	// fmt.Scanf("%f\n", &t)
	// fmt.Println(problem1(r, t))

	//	PROBLEM 2
	// var studentScore int = 80
	// println(problem2(studentScore))

	//	PROBLEM 3
	// var bilangan int
	// fmt.Printf("bilangan : ")
	// fmt.Scanf("%d\n", &bilangan)
	// problem3(bilangan)

	//	PROBLEM 4
	// fmt.Println(problem4(11))
	// fmt.Println(problem4(13))
	// fmt.Println(problem4(17))
	// fmt.Println(problem4(20))
	// fmt.Println(problem4(35))

	//	PROBLEM 6
	fmt.Println(problem6(2, 3))
	fmt.Println(problem6(5, 3))
	fmt.Println(problem6(10, 2))
	fmt.Println(problem6(2, 5))
	fmt.Println(problem6(7, 3))

}

func problem1(r float32, t float32) float32 {
	const phi float32 = 3.14
	result := 2 * phi * float32(r) * (float32(r) + float32(t))

	return result
}

func problem2(grade int) string {

	if grade >= 0 && grade <= 34 {
		return "E"
	} else if grade >= 35 && grade <= 49 {
		return "D"
	} else if grade >= 50 && grade <= 64 {
		return "C"
	} else if grade >= 65 && grade <= 79 {
		return "B"
	} else if grade >= 80 && grade <= 100 {
		return "A"
	} else {
		return "nilai melampaui batas 0 - 100"
	}
}

func problem3(value int) {

	for i := 1; i <= value; i++ {
		if value%i == 0 {
			fmt.Println(i)
		}
	}
}

func problem4(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func problem6(base, pangkat int) int {
	var result int = 1
	for i := 0; i < pangkat; i++ {
		result *= base
	}

	return result
}
