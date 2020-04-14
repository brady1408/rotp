package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/brady1408/rotp/helper"
	"github.com/brady1408/rotp/withstruct"
)

func main() {
	// Setup flags
	genTestFile := flag.Bool("gen-test-file", false, "Generate a test file")
	testFileName := flag.String("test-file-name", "", "Test file name")
	numOfCases := flag.Int("cases", 0, "Specify a number of cases, if not included the cases will be a random number between 1 and 100")
	datasetSize := flag.String("dataset-size", "small", "Can be either \"large\" or \"small\" default is small")
	testString := flag.String("test-string", "", "used to send a string for testing \"++--+--\"")
	printStrings := flag.Bool("print-strings", false, "Used to log the case string with the results")

	flag.Parse()

	if *genTestFile == true {
		helper.GenTestFile(*testFileName, *numOfCases, *datasetSize)
		return
	}

	if *testString != "" {
		flips, err := withstruct.CountFlips(*testString)
		if err != nil {
			fmt.Printf("Error processing %s, Error: %s\n", *testString, err.Error())
		}
		fmt.Printf("Case #%d: %d\n", 1, flips)
		return
	}

	stdin, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if stdin.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("Nothing to process")
		fmt.Println("This command is designed to be used with a pipe")
		fmt.Println("Example: cat filename | rotp")
		fmt.Println("For for help with other options use -h")
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
			// strip \n and maybe \r depending on OS
			s = strings.ReplaceAll(s, "\n", "")
			s = strings.ReplaceAll(s, "\r", "")
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
		var logCaseString string
		if *printStrings {
			logCaseString = s
		}
		fmt.Printf("Case #%d: %d %s\n", i+1, flips, logCaseString)
	}
}
