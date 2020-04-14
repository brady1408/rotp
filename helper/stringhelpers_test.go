package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseHelper(t *testing.T) {
	testString := Reverse("abcdefg")
	assert.Equal(t, "gfedcba", testString)
}

func TestInverseHelper(t *testing.T) {
	testString := Inverse("-+--+-")
	assert.Equal(t, "+-++-+", testString)
}
