package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Due to the random nature of the test file generator our testing is a little limited.
func TestGenerateTestFile(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			"test 42",
			42,
		},
		{
			"test 7",
			7,
		},
		{
			"test 26",
			26,
		},
		{
			"test 87",
			87,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testString := generateCaseString(tt.length)
			assert.Equal(t, tt.length, len(testString))
		})
	}
}
