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
	problem2("abcd")
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

func problem2(angka string) []int {
	var tempMap = map[string]bool{}
	var result []int
	for i := 0; i < len(angka); i++ {
		if tempMap[string(angka[i])] != true {
			tempMap[string(angka[i])] = true
			result = append(result, strconv.ParseInt(string(angka[i])))
		}
	}

	return []int{1, 2, 3}
}
