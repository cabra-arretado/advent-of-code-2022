package main

import (
	"bufio"
	"fmt"
)

func callbackQuestion1(scanner *bufio.Scanner) int {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return 1
}

func answerQ1(file_path string) {
	readFile(file_path, callbackQuestion1)
}
