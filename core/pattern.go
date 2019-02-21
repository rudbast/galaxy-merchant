package core

import (
	"errors"
	"regexp"
	"strings"
)

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

// ExecStatementNumeric parses given text and assign a roman number to given input text.
func ExecStatementNumeric(text string) error {
	if !patternStatementNumeric.MatchString(text) {
		return errors.New("pattern: input doesn't match rule")
	}

	parts := strings.Split(text, " is ")

	if len(parts) != 2 {
		return errors.New("pattern: illegal input")
	}

	numberConversionMap[parts[0]] = rune(parts[1][0])

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
