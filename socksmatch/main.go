package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	matchingSockPairs := 0

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}

	socks := strings.Split(lines[1], " ")
	sockMatch := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		sock, err := strconv.Atoi(socks[i])
		if err != nil {
			panic(err)
		}

		if sockMatch[sock] == true {
			matchingSockPairs += 1
			sockMatch[sock] = false
		} else {
			sockMatch[sock] = true
		}
	}

	fmt.Println(matchingSockPairs)
}
