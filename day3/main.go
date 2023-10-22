package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	var buffer []string
	var counter int


	for scanner.Scan() {
		line := scanner.Text()
		if len(buffer) < 3 {
			buffer = append(buffer, line)
		} 
		if len(buffer) == 3 {
			total += process_line(buffer)
			buffer = []string{}
		}
	}
	fmt.Println("lines >>", counter)
	return total
}

func convert_rune(c rune) int {
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

func process_line(buffer []string) int {
	var total int
	var seen []int
	for _, c := range buffer[0] {
		if strings.ContainsRune(buffer[1], c) && !slices.Contains(seen, int(c)) {
			seen = append(seen, int(c))
		}
	}
	var seen2 []int
	for _, c := range buffer[2] {
		i_c := int(c)
		if slices.Contains(seen, i_c) && !slices.Contains(seen2, i_c){
			seen2 = append(seen2, i_c)
			total += convert_rune(c)
		}
	}
	return total
}

func process_line_q1(l string) int {
	half_index := len(l) / 2
	arr_a := l[:half_index]
	arr_b := l[half_index:]

	var total int
	var seen []int
	for _, c := range arr_a {
		if strings.ContainsRune(arr_b, c) && !slices.Contains(seen, int(c)) {
			total += convert_rune(c)
			seen = append(seen, int(c))
		}
	}
	return total
}

func main() {
	fmt.Println("second answer: ", read_file("./input.txt"))

	// Test question 2
	slice := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}
	slice2 := []string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	fmt.Println(process_line(slice) + process_line(slice2) == 70)
}
