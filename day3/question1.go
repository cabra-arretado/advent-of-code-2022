package main

import (
	"fmt"
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
		total += processLineQ1(line)
	}
	return total
}

func processLineQ1(l string) int {
	half_index := len(l) / 2
	arr_a := l[:half_index]
	arr_b := l[half_index:]

	var total int
	for _, c := range arr_a {
		if strings.ContainsRune(arr_b, c) {
			total += convertRune(c)
		}
	}
	return total
}


func testQ1(){
	// Test cases
	test_1 := processLineQ1("vJrwpWtwJgWrhcsFMMfFFhFp")
	test_2 := processLineQ1("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	test_3 := processLineQ1("PmmdzqPrVvPwwTWBwg")
	fmt.Println("Test case 1 passed:", test_1 == 16)
	fmt.Println("Test case 2 passed", test_2 == 38)
	fmt.Println("Test case 3 passed", test_3 == 42)
}
