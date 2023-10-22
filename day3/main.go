package main

import (
	// "slices"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read_file(file_path string) *bufio.Scanner {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

func process_input(file_path string) int {
	total := 0

	scanner := read_file(file_path)

	for scanner.Scan() {
		line := scanner.Text()
		total += process_line(line)
	}
	return total
}

func convert_rune(c rune) int {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	// Convert rune into the numbers given by the day3
	ci := int(c)

	if ci >= 97 && ci <= 122 {
		return ci - 96
	} else if ci >= 65 && ci <= 90 {
		return ci - 64 + 26
	}
	panic("Did not have a match")
}

func process_line(l string) int {
	half_index := len(l) / 2
	arr_a := l[:half_index]
	arr_b := l[half_index:]

	var total int
	for _, c := range arr_a {
		if strings.ContainsRune(arr_b, c) {
			total += convert_rune(c)
		}
	}
	return total
}

func main() {
	// read_file("./input.txt")

	// Test cases
	// test_1 := process_line("vJrwpWtwJgWrhcsFMMfFFhFp")
	test_2 := process_line("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	test_3 := process_line("PmmdzqPrVvPwwTWBwg")
	// fmt.Println(test_1 == 16)
	fmt.Println(test_2)
	fmt.Println(test_2 == 38)
	fmt.Println(test_3)
	fmt.Println(test_3 == 42)
}
