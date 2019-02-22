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
		testCase{"MMIII", 2003},
		testCase{"LXXXIX", 89},
		testCase{"DCCCXC", 890},
		testCase{"MMMCM", 3900},
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

func TestParseOccurenceExceeded(t *testing.T) {
	cases := []string{
		"IIII",
		"MCCCCX",
		"MDCLLXXXX",
		"MMMMCM",
	}

	for _, c := range cases {
		_, err := Parse(c)
		if err != ErrCharOcurrenceLimit {
			t.Errorf("Parse failure, input: %s, expected error: %s, got: %s.", ErrCharOcurrenceLimit, err, c)
			return
		}
	}
}
