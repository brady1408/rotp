5
-
-+
+-
+++
--+-
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/brady1408/rotp/withstruct"
)

func main() {
	stdin, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if stdin.Size() == 0 {
		fmt.Println("Nothing to process")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var input []string

	for {
		s, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		if s != "" {
			input = append(input, s)
		}
	}

	// first input should be the number of cases
	// cases should be between 1 and 100
	cases, err := strconv.Atoi(input[0])
	if err != nil || cases < 1 || cases > 100 {
		fmt.Println("The first line on input should be a number between 1 and 100")
	}

	// now that we have cases lets strip it off
	// everything left should be a test string
	input = input[1:]

	// sinse they gave us the number of cases lets make sure that they match
	if cases != len(input) {
		fmt.Println("The number of cases provided does not match the number of datasets")
	}

	for i, s := range input {
		flips, err := withstruct.CountFlips(s)
		if err != nil {
			fmt.Printf("Error processing %s, Error: %s\n", s, err.Error())
		}
		fmt.Printf("Case #%d: %d\n", i, flips)
	}
}
