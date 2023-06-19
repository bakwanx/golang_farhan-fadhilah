package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5, a6 int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	var min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min : ", min)
	fmt.Println("Nilai Max : ", max)
}

func getMinMax(numbers ...*int) (min, max int) {
	var minTemp = *numbers[0]
	var maxTemp = *numbers[0]
	for i, _ := range numbers {
		if *numbers[i] < minTemp {
			minTemp = *numbers[i]
		}
		if *numbers[i] > maxTemp {
			maxTemp = *numbers[i]
		}

	}
	return minTemp, maxTemp
}
