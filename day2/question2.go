package main

import (
	"bufio"
	"os"
	"strings"
)

func readFileQ2(file_path string) int {
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
		total += processLineQ2(line)
	}
	return total
}

func processLineQ2(l string) int {
	// For Question 2
	// A: Rock : 1 point
	// B: Paper : 2 points
	// C: Scissors : 3 points
	// X: 0 points for loss
	// Y: 3 points for draw
	// Z: 6 points for win
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
