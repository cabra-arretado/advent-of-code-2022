package main

import (
	"bufio"
	"os"
	"strings"
)

func readFileQ1(file_path string) int {
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

func processLineQ1(l string) int {
	// A: X: Rock : 1 point
	// B: Y: Paper : 2 points
	// C: Z: Scissors : 3 points
	// 0 points for loss
	// 3 points for draw
	// 6 points for win
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
