package main

import (
	"slices"
	"bufio"
	"fmt"
	"os"
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

func process_line(l string) int {
	rune_arr := []rune(l)
	half_index := len(rune_arr) / 2
	arr_a := rune_arr[:half_index]
	arr_b := rune_arr[half_index:]

	for e := range arr_a{
		slices.Contains(arr_b, e)
	}
		
	return 0
}

func main() {
	a := process_line("vJrwpWtwJgWrhcsFMMfFFhFp")
	b := process_line("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	c := process_line("PmmdzqPrVvPwwTWBwg")
	fmt.Println(a == 16)
	fmt.Println(b == 38)
	fmt.Println(c == 42)
}
