package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	text := strings.Split(string(content), ",")
	var s []int

	for i := 0; i < len(text); i++ {
		num, err := strconv.Atoi(text[i])
		if err != nil {
			fmt.Println(err)
		}
		s = append(s, num)
	}

	s[1] = 12
	s[2] = 2

	for opcode := 0; opcode < len(s); opcode +=4 {
		switch s[opcode] {
		case 1:
			s[s[opcode + 3]] = s[s[opcode + 1]] + s[s[opcode + 2]]
		case 2:
			s[s[opcode + 3]] = s[s[opcode + 1]] * s[s[opcode + 2]]
		case 99:
			break
		default:
			fmt.Println("Opcode is neither 1,2 nor 99")
		}
	}

	fmt.Println(s[0])
}