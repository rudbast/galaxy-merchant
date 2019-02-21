package core

import (
	"fmt"
	"testing"
)

func TestExecStatementNumeric(t *testing.T) {
	type testCase struct {
		Input  string
		Output rune
	}

	cases := []testCase{
		testCase{"glob", 'I'},
		testCase{"prok", 'V'},
		testCase{"pish", 'X'},
		testCase{"tegj", 'L'},
	}

	for _, c := range cases {
		input := fmt.Sprintf("%s is %c", c.Input, c.Output)

		err := ExecStatementNumeric(input)
		if err != nil {
			t.Errorf("Execute error for input: %s, error: %v\n", c.Input, err)
			return
		}

		if val := numberConversionMap[c.Input]; val != c.Output {
			t.Errorf("Execute failure, input: %s, expected: %c, got: %c.", input, c.Output, val)
			return
		}
	}
}
