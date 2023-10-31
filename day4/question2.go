package main

import (
	"bufio"
	"fmt"
	"strings"
)

func processLineQ2(s string) int {
	var aBegin, aEnd, bBegin, bEnd int
	_, err := fmt.Sscanf(s, "%d-%d,%d-%d", &aBegin, &aEnd, &bBegin, &bEnd)
	if err != nil {
		panic(err)
	}
	// If a touches b
	//TODO: Correct this logic
	if aBegin <= bBegin && aEnd >= bEnd {
		return 1
	}
	// If b touches a
	if bBegin <= aBegin && bEnd >= aEnd {
		return 1
	}
	return 0
}

func callbackQuestion2(scanner *bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += processLineQ2(line)
	}
	return total
}

func answerQ2(file_path string) {
	fmt.Println("Question 2:", readFile(file_path, callbackQuestion2))
}

func testQuestion2() {
	arr := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	testLines := strings.Split(arr, "\n") 
	expected := []int{0, 0, 1, 1, 1, 1}
	for i := range testLines{
		fmt.Println(processLineQ2(testLines[i]) == expected[i])
	}
}
