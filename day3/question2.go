package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	var buffer []string

	for scanner.Scan() {
		line := scanner.Text()
		if len(buffer) < 3 {
			buffer = append(buffer, line)
		}
		if len(buffer) == 3 {
			total += processLineQ2(buffer)
			buffer = []string{}
		}
	}
	return total
}

func processLineQ2(buffer []string) int {
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
		if slices.Contains(seen, i_c) && !slices.Contains(seen2, i_c) {
			seen2 = append(seen2, i_c)
			total += convertRune(c)
		}
	}
	return total
}

func testQ2() {
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
	fmt.Println("Test case passed:", processLineQ2(slice)+processLineQ2(slice2) == 70)
}
