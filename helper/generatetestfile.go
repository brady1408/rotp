package helper

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"strings"
)

// GenTestFile will write a test file to disk to use with this application
func GenTestFile(testFileName string, numOfCases int, datasetSize string) {
	// validate data
	if testFileName == "" {
		fmt.Println("file name cannot be blank")
		return
	}
	if numOfCases > 100 {
		fmt.Println("Number of cases needs to be a number between 1 and 100")
		return
	}
	if strings.ToLower(datasetSize) != "large" && strings.ToLower(datasetSize) != "small" {
		fmt.Println("Dataset size needs to be set to either large or small")
		return
	}

	size := 100
	if strings.ToLower(datasetSize) == "small" {
		size = 10
	}

	if numOfCases == 0 {
		numOfCases = rand.Intn(100)
	}

	output := fmt.Sprintf("%d\n", numOfCases)

	//Generate cases
	for i := 0; i < numOfCases; i++ {
		l := rand.Intn(size)
		output += fmt.Sprintf("%s\n", generateCaseString(l))
	}

	err := ioutil.WriteFile(testFileName, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func generateCaseString(length int) string {
	var caseString string
	for i := 0; i < length; i++ {
		s := "-"
		if math.Mod(float64(rand.Intn(100)), 2) == 0 {
			s = "+"
		}
		caseString += s
	}
	return caseString
}
