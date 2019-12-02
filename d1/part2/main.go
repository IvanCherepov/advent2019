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

    for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		for number/3 -2 > 0 {
			result += number/3 -2 
			number = number/3 -2
		}	
	}

	fmt.Println(result)
}

