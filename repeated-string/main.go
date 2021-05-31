package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func numChar(s string, c byte, max int) int {
	var count int

	for i := range s {
		if i > max-1 {
			break
		}
		if s[i] == c {
			count++
		}
	}

	return count
}

func main() {
	var repeatString string

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		repeatString = input.Text()
	}

	var n int64
	if input.Scan() {
		n, _ = strconv.ParseInt(input.Text(), 10, 64)
	}

	if err := input.Err(); err != nil {
		panic("unable to parse input")
	}

	var count int
	length := len(repeatString)
	repeatNum := (n / int64(length))
	charsInString := numChar(repeatString, 'a', length)
	remainingPositions := int(n % int64(length))
	remainingChars := int64(numChar(repeatString, 'a', remainingPositions))

	count = (int(repeatNum) * charsInString) + int(remainingChars)

	fmt.Println(count)
}
