package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Min() (min int, name string) {
	var minTemp = s.score[0]
	var nameMinTemp = s.name[0]

	for i, v := range s.score {
		if v < minTemp {
			minTemp = v
			nameMinTemp = s.name[i]
		}
	}
	return minTemp, nameMinTemp
}

func (s Student) Max() (max int, name string) {
	var maxTemp = s.score[0]
	var nameMaxTemp = s.name[0]
	for i, v := range s.score {
		if v > maxTemp {
			maxTemp = v
			nameMaxTemp = s.name[i]
		}
	}
	return maxTemp, nameMaxTemp
}

func (s Student) Average() float64 {
	var length = len(s.score)
	var sumTemp float64 = 0.0
	for _, v := range s.score {
		sumTemp += float64(v)
	}
	return sumTemp / float64(length)
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input : " + string(1) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input : " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is: ", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is: ", scoreMax, nameMax)
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is: ", scoreMin, nameMin)
}
