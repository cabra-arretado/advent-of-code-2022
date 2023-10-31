package main

import (
	"bufio"
	"fmt"
)

func processLineQ1(s string) int {
	var aBegin, aEnd, bBegin, bEnd int
	_, err := fmt.Sscanf(s, "%d-%d,%d-%d", &aBegin, &aEnd, &bBegin, &bEnd)
	if err != nil {
		panic(err)
	}
	// If A contains B
	if aBegin <= bBegin && aEnd >= bEnd {
		return 1
	}
	// If B contains A
	if bBegin <= aBegin && bEnd >= aEnd {
		return 1
	}
return 0
}


func callbackQuestion1(scanner *bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += processLineQ1(line)
	}
	return total
}

func answerQ1(file_path string) {
	fmt.Println("Question 1:", readFile(file_path, callbackQuestion1))
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
