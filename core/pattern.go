package core

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/rudbast/galaxy-merchant/roman"
)

var (
	// Input pattern rules.
	patternStatementNumeric = regexp.MustCompile("\\w+\\sis\\s(I|V|X|L|C|D|M)")
	patternStatementValue   = regexp.MustCompile(".+\\s\\w+\\sis\\s\\d+\\sCredits")
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

// ExecStatementValue parses given text & computes the value per unit value of given material count.
// Example:
// glob glob Silver is 34 Credits
// glob prok Gold is 57800 Credits
func ExecStatementValue(text string) error {
	if !patternStatementValue.MatchString(text) {
		return errors.New("pattern: input doesn't match rule")
	}

	parts := strings.Split(text, " is ")

	if len(parts) != 2 {
		return errors.New("pattern: invalid input length")
	}

	numParts := strings.Split(parts[0], " ")
	material := numParts[len(numParts)-1]

	count, err := extractMaterialCount(numParts[:len(numParts)-1])
	if err != nil {
		return errors.Wrap(err, "pattern")
	}

	credit, err := strconv.ParseFloat(strings.TrimSuffix(parts[1], " Credits"), 64)
	if err != nil {
		return errors.Wrap(err, "pattern")
	}

	// Material per unit value.
	unit := credit / float64(count)
	materialUnitMap[material] = unit

	return nil
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

func ExecQuestionNumeric(text string) (string, error) {
	return "", nil
}

func ExecQuestionValue(text string) (string, error) {
	return "", nil
}
