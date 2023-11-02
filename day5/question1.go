package main

import (
	"bufio"
	"fmt"
	// "strings"

)
/*
The first 9 lines of the file contain the curent state.
The matrix is 9x35 containig trailing spaces int the end if needed

    [B]             [B] [S]        
    [M]             [P] [L] [B] [J]
    [D]     [R]     [V] [D] [Q] [D]
    [T] [R] [Z]     [H] [H] [G] [C]
    [P] [W] [J] [B] [J] [F] [J] [S]
[N] [S] [Z] [V] [M] [N] [Z] [F] [M]
[W] [Z] [H] [D] [H] [G] [Q] [S] [W]
[B] [L] [Q] [W] [S] [L] [J] [W] [Z]
 1   2   3   4   5   6   7   8   9 
*/

func processLineQ1(s string) int {
	return 0
}

func callbackQ1(scanner *bufio.Scanner) int {
	return 0
}

func answerQ1(file_path string) {
	fmt.Println("Question 1:", readFile(file_path, callbackQ1))
}

func testQ1() {
}
