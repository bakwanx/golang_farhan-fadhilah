package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*-------Problem 1-------*/
	// problem1([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"})
	// problem1([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"})
	// problem1([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"})

	/*-------Problem 2-------*/
	// fmt.Println(problem2("1234123"))
	// fmt.Println(problem2("76523752"))
	// fmt.Println(problem2("12345333"))

	/*-------Problem 3-------*/
	fmt.Println(problem3([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(problem3([]int{2, 5, 9, 11}, 11))
	fmt.Println(problem3([]int{1, 3, 5, 7}, 12))
}

func problem1(arrayA []string, arrayB []string) []string {

	var joinData = append(arrayA, arrayB...)
	var tempMap = map[string]bool{}
	var result []string

	for i := range joinData {

		if tempMap[joinData[i]] != true {
			tempMap[joinData[i]] = true
			result = append(result, joinData[i])
		}
	}
	for _, e := range result {
		fmt.Println(e)
	}
	fmt.Println("/*------------------*/")
	return result
}

func problem2(angka string) []int64 {

	var result []int64
	for i := 0; i < len(angka); i++ {
		sum := 0
		var num, _ = strconv.ParseInt(string(angka[i]), 10, 0)
		for j := 0; j < len(angka); j++ {
			var subNum, _ = strconv.ParseInt(string(angka[j]), 10, 0)
			if num == subNum {
				sum++
			}
		}

		if sum <= 1 {
			result = append(result, num)
		}

	}

	return result
}

func problem3(arr []int, target int) []int64 {
	var result []int64
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			var tempSum = arr[i] + arr[j]
			if tempSum == target && arr[i] != arr[j] {
				result = append(result, int64(i))
			}
		}
	}

	return result
}
