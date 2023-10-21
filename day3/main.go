package main

import (
	"slices"
	"bufio"
	"fmt"
	"os"
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

func process_line(l string) int {
	rune_arr := []rune(l)
	half_index := len(rune_arr) / 2
	arr_a := rune_arr[:half_index]
	arr_b := rune_arr[half_index:]

	for e := range arr_a{
		if slices.Contains(arr_b, e){
			return e
		}
	}
		
	return 0
}

func main() {
	// read_file("./input.txt")


	// Test cases
	test_1 := process_line("vJrwpWtwJgWrhcsFMMfFFhFp")
	test_2 := process_line("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	test_3 := process_line("PmmdzqPrVvPwwTWBwg")
	fmt.Println(test_1 == 16)
	fmt.Println(test_2 == 38)
	fmt.Println(test_3 == 42)
}
