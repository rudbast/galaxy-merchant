package core

import "regexp"

var (
	// Input pattern rules.
	patternStatementNumeric = regexp.MustCompile("\\w+\\sis\\s(I|V|X|L|C|D|M)")
	patternStatementValue   = regexp.MustCompile(".+\\sis\\s\\d+\\sCredits")
	patternQuestionNumeric  = regexp.MustCompile("how much is .+\\s\\?")
	patternQuestionValue    = regexp.MustCompile("how many Credits is .+\\s\\?")
)

var (
	// Numeral conversion mapper from input text to roman number.
	numberConversionMap = map[string]rune{}

	// Material value per unit map.
	materialUnitMap = map[string]float64{}
)

func ExecStatementNumeric(text string) error {
	return nil
}

func ExecStatementValue(text string) error {
	return nil
}

func ExecQuestionNumeric(text string) (string, error) {
	return "", nil
}

func ExecQuestionValue(text string) (string, error) {
	return "", nil
}
