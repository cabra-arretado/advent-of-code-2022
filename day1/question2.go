package main

import (
	"bufio"
	"os"
	"strconv"
	"sort"
)

func secondQuestion(file_path string) []int {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	list := []int{}
	totalCal := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			cal, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			totalCal += cal
			continue
		}
		list = append(list, totalCal)
		totalCal = 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	return list[:3]
}
