package main

import (
	"bufio"
	"os"
	"strconv"
)

func firstQuestion(file_path string) int {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalCal := 0
	maxCal := 0

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
		if totalCal > maxCal {
			maxCal = totalCal
		}
		totalCal = 0
	}
	return maxCal
}
