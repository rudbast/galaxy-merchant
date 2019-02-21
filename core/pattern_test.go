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
			t.Errorf("Execute error for input: %s, error: %v\n", input, err)
			return
		}

		if val := numberConversionMap[c.Input]; val != c.Output {
			t.Errorf("Execute failure, input: %s, expected: %c, got: %c.", input, c.Output, val)
			return
		}
	}
}

func TestExecStatementValue(t *testing.T) {
	type testCase struct {
		InputCount    string
		InputMaterial string
		OutputTotal   float64
		OutputUnit    float64
	}

	cases := []testCase{
		testCase{"glob glob", "Silver", 34, 17},
		testCase{"glob prok", "Gold", 57800, 14450},
		testCase{"pish pish", "Iron", 3910, 195.5},
	}

	for _, c := range cases {
		input := fmt.Sprintf("%s %s is %.0f Credits", c.InputCount, c.InputMaterial, c.OutputTotal)

		err := ExecStatementValue(input)
		if err != nil {
			t.Errorf("Execute error for input: %s, error: %v\n", input, err)
			return
		}

		if val := materialUnitMap[c.InputMaterial]; val != c.OutputUnit {
			t.Errorf("Execute failure, input: %s, expected: %f, got: %f.", input, c.OutputUnit, val)
			return
		}
	}
}

func TestExecQuestionNumeric(t *testing.T) {
	type testCase struct {
		Input  string
		Output int64
	}

	cases := []testCase{
		testCase{"pish tegj glob glob", 42},
	}

	for _, c := range cases {
		input := fmt.Sprintf("how much is %s ?", c.Input)
		output := fmt.Sprintf("pish tegj glob glob is %d", c.Output)

		result, err := ExecQuestionNumeric(input)
		if err != nil {
			t.Errorf("Execute error for input: %s, error: %v\n", input, err)
			return
		}

		if result != output {
			t.Errorf("Execute failure, input: %s, expected: %s, got: %s.", input, output, result)
			return
		}
	}
}

func TestExecQuestionValue(t *testing.T) {
	type testCase struct {
		InputCount    string
		InputMaterial string
		Output        float64
	}

	cases := []testCase{
		testCase{"glob prok", "Silver", 68},
		testCase{"glob prok", "Gold", 57800},
		testCase{"glob prok", "Iron", 782},
	}

	for _, c := range cases {
		input := fmt.Sprintf("how many Credits is %s %s ?", c.InputCount, c.InputMaterial)
		output := fmt.Sprintf("%s %s is %.0f Credits", c.InputCount, c.InputMaterial, c.Output)

		result, err := ExecQuestionValue(input)
		if err != nil {
			t.Errorf("Execute error for input: %s, error: %v\n", input, err)
			return
		}

		if result != output {
			t.Errorf("Execute failure, input: %s, expected: %s, got: %s.", input, output, result)
			return
		}
	}
}
