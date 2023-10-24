package main

import "fmt"

func convertRune(c rune) int {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	// Convert rune into the numbers given by the day3 q1
	ci := int(c)

	if ci >= 97 && ci <= 122 {
		return ci - 96
	} else if ci >= 65 && ci <= 90 {
		return ci - 38
	}
	panic("Did not have a match")
}

func main(){
	file_path := "./input.txt"
	fmt.Println("First answer:", readFileQ1(file_path))
	fmt.Println("Second answer:", readFileQ2(file_path))
}
