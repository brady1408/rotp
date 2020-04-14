# rotp (Revenge of the Pancakes)

## Welcome to rotp

This was a fun exercise and I enjoyed working on it. So thank you!

### Usage

This application was designed around using the stdin to pipe test files directly to the app. The exercise talked about "Lines of input" I interpreted this to mean that it would take a multi line input. Stdin seems like the natural way to accomplish this.

### Command Help

Running rotp without any input will give you a little bit of help

    Nothing to process
	This command is designed to be used with a pipe
	Example: cat filename | rotp
	For for help with other options use -h
More help can be found by using the command with the -h flag

	rotp -h
	
### Flags

 - -gen-test-file - this will cause the app to output a randomly generated test file that you can then pipe back in the the app for testing.
 - -test-file-name - The file name for the test file being generated
 - -cases - The number of cases to include in the test app. This needs to be a number between 1 and 100 according to the excersise documentation
 - -dataset-size - The test file can be generated with either Large or Small datasets. This parameter will accept either small or large as an argument. small will produce test cases with a length of 1 to 10 large will produce test cases with a length of 1 to 100
 - -print-strings - The normal output of the app is a list formatted as
 
		Case #1: 1 
		Case #2: 1 
		Case #3: 2 
		Case #4: 0 
		Case #5: 3 
	the -print-strings flag will include the string that was tested. changing the output to be
	
		Case #1: 1 -
		Case #2: 1 -+
		Case #3: 2 +-
		Case #4: 0 +++
		Case #5: 3 --+-

- -test-string will allow you to send the app a single string for testing, this needs to be formatted as a stack specified in the documentation

		+ for face - for blank between 1 and 100 characters in length

### Examples

Generate a test file with 10 cases and small datasets.

	rotp -test-file-name first.txt -gen-test-file -cases 10 -dataset-size small

processing a file with test cases.

	cat first.txt | rotp

processing a file with test cases with the case string included in the output

	cat first.txt | rotp -print-strings

sending a single string to the app for testing

	rotp -test-string +--+--+



### Compiling

	go build

or

	go install

the app can also be run with go run

	go run main.go -test-file-name first.txt -gen-test-file -cases 10 -dataset-size small
	
	cat first.txt | go run main.go

### Tests

Tests can be run with the go test tool

	go test ./...

