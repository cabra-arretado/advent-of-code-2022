package main

import (
	"fmt"
)

func main() {
	file := "./input.txt"
	answer1 := firstQuestion(file)
	fmt.Println("First question:", answer1)
	total := 0 
	answer2 := secondQuestion(file)
	for v := range answer2 {
		total += answer2[v]
	}
	fmt.Println("second question:", total)
}
