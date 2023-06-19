package main

import "strings"

func main() {

	println(problem1("AKA", "AKASHI"))
	println(problem1("KANGOORO", "KANG"))
}

func problem1(a, b string) string {
	result := strings.Contains(b, a)
	if !result {
		return "tidak ada kata" + a
	}
	return a
}
