package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

    scanner := bufio.NewScanner(file)

	result := 0

    for scanner.Scan() {             // internally, it advances token based on sperator
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result += number/3 - 2
	}
	
	fmt.Println(result)
}

