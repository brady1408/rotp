package withstruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := stack{
		value: "+-+-+",
	}
	top := s.findTop()
	assert.Equal(t, 1, top)
	bottom := s.findBottom()
	assert.Equal(t, 4, bottom)
	s.flip(top)
	assert.Equal(t, "--+-+", s.value)
	s.flip(bottom)
	assert.Equal(t, "+-+++", s.value)
}

func TestCases(t *testing.T) {
	tests := []struct {
		name    string
		stack   string
		want    int
		wantErr bool
	}{
		{
			"test -",
			"-",
			1,
			false,
		},
		{
			"test -+",
			"-+",
			1,
			false,
		},
		{
			"test +-",
			"+-",
			2,
			false,
		},
		{
			"test +++",
			"+++",
			0,
			false,
		},
		{
			"test --+-",
			"--+-",
			3,
			false,
		},
		{
			"Test blank",
			"",
			0,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := CountFlips(tt.stack)
			// This is kind of ugly, maybe a if else is better even...
			if (err != nil) != tt.wantErr {
				t.Errorf("Error was not nil: error = %w, WantErr %t", err, tt.wantErr)
			}
			assert.Equal(t, tt.want, f)
		})
	}
}
