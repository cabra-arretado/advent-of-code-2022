package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_file(file_path string) int {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total_cal := 0
	max_cal := 0

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
		if total_cal > max_cal {
			fmt.Println("max_calories")
			fmt.Println(max_cal)
			fmt.Println("total_calories")
			fmt.Println(total_cal)
			max_cal = total_cal
		}
		total_cal = 0
	}

	return total_cal

}

func main() {
	file := "./input.txt"
	answer := read_file(file)
	fmt.Println(answer)
}
