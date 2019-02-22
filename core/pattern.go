package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/rudbast/galaxy-merchant/roman"
)

var (
	// Input pattern rules.
	patternStatementNumeric = regexp.MustCompile(`^(\w+) is (I|V|X|L|C|D|M)$`)
	patternStatementValue   = regexp.MustCompile(`^(.+) (\w+) is ([0-9]*\.?[0-9]+) [Cc]redits$`)
	patternQuestionNumeric  = regexp.MustCompile(`^[Hh]ow (much|many) is (.+) \?$`)
	patternQuestionValue    = regexp.MustCompile(`^[Hh]ow (much|many) [Cc]redits is (.+) (\w+) \?$`)

	// Numeral conversion mapper from input text to roman number.
	numberConversionMap = map[string]rune{}

	// Material value per unit map.
	materialUnitMap = map[string]float64{}

	ErrPatternMismatch = errors.New("pattern: input doesn't match rule")
)

// ExecStatementNumeric parses given text and assign a roman number to given input text.
// Example:
// - glob is I
// - prok is V
func ExecStatementNumeric(text string) (string, error) {
	if !patternStatementNumeric.MatchString(text) {
		return "", ErrPatternMismatch
	}

	parts := patternStatementNumeric.FindStringSubmatch(text)
	word := parts[1]
	romanNum := parts[2][0]

	numberConversionMap[word] = rune(romanNum)

	return "", nil
}

// ExecStatementValue parses given text & computes the value per unit value of given material count.
// Example:
// glob glob Silver is 34 Credits
// glob prok Gold is 57800 Credits
func ExecStatementValue(text string) (string, error) {
	if !patternStatementValue.MatchString(text) {
		return "", ErrPatternMismatch
	}

	parts := patternStatementValue.FindStringSubmatch(text)
	numParts := strings.Split(parts[1], " ")
	material := parts[2]
	creditStr := parts[3]

	count, err := extractMaterialCount(numParts)
	if err != nil {
		return "", errors.Wrap(err, "pattern: statement value")
	}

	credit, err := strconv.ParseFloat(creditStr, 64)
	if err != nil {
		return "", errors.Wrap(err, "pattern")
	}

	// Material per unit value.
	unit := credit / float64(count)
	materialUnitMap[material] = unit

	return "", nil
}

// extractMaterialCount parses decimal values of given roman numeral word.
func extractMaterialCount(words []string) (int64, error) {
	var romanNum string

	// Convert / parse given number word into roman numeral.
	for _, word := range words {
		num, ok := numberConversionMap[word]
		if !ok {
			return 0, errors.Errorf("extract: unknown number word %s", word)
		}

		romanNum += string(num)
	}

	val, err := roman.Parse(romanNum)
	if err != nil {
		return 0, errors.Wrap(err, "extract")
	}

	return val, nil
}

// ExecQuestionNumeric parses given text & returns the decimal value of given roman numeral word.
func ExecQuestionNumeric(text string) (string, error) {
	if !patternQuestionNumeric.MatchString(text) {
		return "", ErrPatternMismatch
	}

	parts := patternQuestionNumeric.FindStringSubmatch(text)
	words := parts[2]

	count, err := extractMaterialCount(strings.Split(words, " "))
	if err != nil {
		return "", errors.Wrap(err, "pattern: question numeric")
	}

	return fmt.Sprintf("%s is %d", words, count), nil
}

// ExecQuestionValue parses given text & returns the decimal value of given material & roman numeral word.
func ExecQuestionValue(text string) (string, error) {
	if !patternQuestionValue.MatchString(text) {
		return "", ErrPatternMismatch
	}

	parts := patternQuestionValue.FindStringSubmatch(text)
	words := parts[2]
	material := parts[3]

	count, err := extractMaterialCount(strings.Split(words, " "))
	if err != nil {
		return "", errors.Wrap(err, "pattern: question value")
	}

	unit, ok := materialUnitMap[material]
	if !ok {
		return "", errors.New("pattern: unknown material")
	}

	value := float64(count) * unit

	return fmt.Sprintf("%s %s is %.0f Credits", words, material, value), nil
}
