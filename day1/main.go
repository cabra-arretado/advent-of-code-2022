package main

import (
	"fmt"
	"os"
)

func read_file(file string) {
	dt, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Print(dt)
}

func main() {
	file := "./input.txt"
	read_file(file)
}
