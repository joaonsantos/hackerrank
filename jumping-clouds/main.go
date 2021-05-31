package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const maxSteps int = 100
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	// get total number of clouds
	input.Scan()
	numClouds, _ := strconv.Atoi(input.Text())

	// fill cloud array, 1 is thunderhead, 0 is cummulus
	clouds := make([]int, numClouds)
	for i := 0; input.Scan(); i++ {
		clouds[i], _ = strconv.Atoi(input.Text())
	}

	if err := input.Err(); err != nil {
		panic("could not parse input")
	}

	var curr int
	var numSteps int
	//var steps [maxSteps]int
	for i := 0; i < numClouds; i++ {
		// in each iteration try to step 2, then 1 if possible
		// save the current step
		for step := 2; step >= 1; step-- {
			if curr+step < numClouds {
				if clouds[curr+step] != 1 {
					// steps[i] = curr
					curr = curr + step
					numSteps++
					break
				}
			}
		}
	}

	fmt.Println(numSteps)
}
