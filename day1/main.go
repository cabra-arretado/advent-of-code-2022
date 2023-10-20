package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func read_file(file_path string) []int {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	list := []int{}
	total_cal := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			cal, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			total_cal += cal
			continue
		}
		list = append(list, total_cal)
		total_cal = 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	return list[:3]
}

func main() {
	file := "./input.txt"
	answer := read_file(file)
	fmt.Println("first question:", answer[0])
	total := 0 
	for v := range answer {
		total += answer[v]
	}
	fmt.Println("second question:", total)
}
