package main

import (
	"bufio"
	"fmt"
	"strings"
)

func processLineQ1(s string) int {
	arr := strings.Split(s, ",")
	if contains(arr[0], arr[1]) {
		return 1
	}
	return 0
}

func contains(s1 string, s2 string) bool {
	arrA := strings.Split(s1, "-")
	arrB := strings.Split(s2, "-")
	// Check the left side first
	if arrA[0] <= arrB[0] && arrA[1] >= arrB[1] {
		return true
	}
	if arrB[0] <= arrA[0] && arrB[1] >= arrA[1] {
		return true
	}
	return false
}

func callbackQuestion1(scanner *bufio.Scanner) int {
	total := 0
	lines := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += processLineQ1(line)
		lines++
	}
	fmt.Println(lines)
	return total
}

func answerQ1(file_path string) {
	fmt.Println(readFile(file_path, callbackQuestion1))
}

func testQuestion1() {
	test1 := "2-4,6-8" // Flase
	test2 := "6-6,4-6" // True
	test3 := "2-8,3-7" // True
	test4 := "2-3,4-5" // False
	test5 := "5-7,7-9" // True
	test6 := "2-6,4-8" // True
	fmt.Println("test", processLineQ1("24-50,25-50"))
	fmt.Println("Test 1:", processLineQ1(test1) == 0)
	fmt.Println("Test 2:", processLineQ1(test2) == 1)
	fmt.Println("Test 3:", processLineQ1(test3) == 1)
	fmt.Println("Test 4:", processLineQ1(test4) == 0)
	fmt.Println("Test 5:", processLineQ1(test5) == 0)
	fmt.Println("Test 6:", processLineQ1(test6) == 0)
}
