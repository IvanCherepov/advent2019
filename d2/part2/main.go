package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func calc(s []int, n int, v int) (int, int) {
	s[1] = n
	s[2] = v
	for opcode := 0; opcode < len(s); opcode += 4 {
		switch s[opcode] {
		case 1:
			s[s[opcode+3]] = s[s[opcode+1]] + s[s[opcode+2]]
		case 2:
			s[s[opcode+3]] = s[s[opcode+1]] * s[s[opcode+2]]
		case 99:
			//fmt.Println("YYYYYYYYYYYY", n, v, s[0], "YYYYYYYYYYYY")
			break
		default:
			//fmt.Println("Opcode is neither 1,2 nor 99")
		}
	}
	if s[0] == 19690720 {
		fmt.Println(100*n + v)
		return n, v
	}

	//fmt.Println(n, v, s[0])
	return 0, 0
}

func strtoint(t []string) []int {
	var s []int
	for i := 0; i < len(t); i++ {
		num, err := strconv.Atoi(t[i])
		if err != nil {
			fmt.Println(err)
		}
		s = append(s, num)
	}
	return s
}

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	text := strings.Split(string(content), ",")

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			calc(strtoint(text), noun, verb)
			//fmt.Println(noun, verb)
		}
	}
}
