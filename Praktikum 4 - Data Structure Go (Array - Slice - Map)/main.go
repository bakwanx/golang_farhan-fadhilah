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

func problem2(angka string) []int64 {
	var temp int64

	for i := 0; i < len(angka); i++ {
		var intCon, _ = strconv.ParseInt(angka[i], 2, 0)
		temp = temp ^ intCon
	}

	return []int64{1, 2, 3}
}
