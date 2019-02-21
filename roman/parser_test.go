package roman

import "testing"

func TestParse(t *testing.T) {
	type testCase struct {
		Input  string
		Output int64
	}

	cases := []testCase{
		testCase{"MMVI", 2006},
		testCase{"MCMXLIV", 1944},
		testCase{"MCMIII", 1903},
	}

	for _, c := range cases {
		val, err := Parse(c.Input)
		if err != nil {
			t.Errorf("Parse error for input: %s, error: %v\n", c.Input, err)
			return
		}

		if val != c.Output {
			t.Errorf("Parse failure, input: %s, expected: %d, got: %d.", c.Input, c.Output, val)
			return
		}
	}
}
