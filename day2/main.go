package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read_file(file_path string) int {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		total += process_line(line)
	}
	return total
}

func process_line(l string) int{
	s := strings.Split(l, " ")
	sum := 0
	switch s[1] {
	case "X": //A // To win B // To lose C
		switch s[0] {
		case "A":
			sum += 3
		case "B":
			sum += 1
		case "C":
			sum += 2
		}
	case "Y": //B // To win C // To lose A
		sum += 3
		switch s[0] {
		case "A":
			sum += 1
		case "B":
			sum += 2
		case "C":
			sum += 3
		}
	case "Z": // C // To win A // To lose B
		sum += 6
		switch s[0] {
		case "A":
			sum += 2
		case "B":
			sum += 3
		case "C":
			sum += 1
		}
	}
	return sum
}
// For Question 2
// A: Rock : 1 point
// B: Paper : 2 points
// C: Scissors : 3 points
// X: 0 points for loss
// Y: 3 points for draw
// Z: points for win
func main() {
	fmt.Println(read_file("./input.txt"))
	// Tests Question 2
	// a := process_line("A Y")
	// b := process_line("B X")
	// fmt.Println(a == 4)
	// fmt.Println(b == 1)
}

/*
// Question 1
// A: X: Rock : 1 point
// B: Y: Paper : 2 points
// C: Z: Scissors : 3 points
// 0 points for loss
// 3 points for draw
// 6 points for win
func process_line(l string) int {
	s := strings.Split(l, " ")
	sum := 0
	switch s[1] {
	case "X":
		sum += 1
		switch s[0] {
		case "A":
			sum += 3
		case "B":
			sum += 0
		case "C":
			sum += 6
		}
	case "Y":
		sum += 2
		switch s[0] {
		case "A":
			sum += 6
		case "B":
			sum += 3
		case "C":
			sum += 0
		}
	case "Z":
		sum += 3
		switch s[0] {
		case "A":
			sum += 0
		case "B":
			sum += 6
		case "C":
			sum += 3
		}
	}
	return sum
}
*/

