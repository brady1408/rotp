package withstruct

import (
	"fmt"
	"strings"

	"github.com/brady1408/rotp/helper"
)

type stack struct {
	flips int
	value string
}

// CountFlips returns the number of flips needed to sort the stack of pancakes
func CountFlips(stackString string) (int, error) {
	if stackString == "" {
		return 0, fmt.Errorf("String cannot be blank")
	}
	s := stack{
		value: stackString,
	}
	for {
		top := s.findTop()
		if top == -1 {
			return s.flips, nil
		}

		bottom := s.findBottom()
		if top > 0 {
			s.flip(top)
		}
		s.flip(bottom)
	}
}

func (s stack) findBottom() int {
	return len(s.value) - strings.Index(helper.Reverse(s.value), "-")
}

func (s stack) findTop() int {
	return strings.Index(s.value, "-")
}

func (s *stack) flip(position int) {
	s.flips++
	s.value = helper.Reverse(helper.Inverse(s.value[:position])) + s.value[position:]
}
