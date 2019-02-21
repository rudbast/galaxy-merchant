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
// Example:
// - glob is I
// - prok is V
func ExecStatementNumeric(text string) error {
	if !patternStatementNumeric.MatchString(text) {
		return errors.New("pattern: input doesn't match rule")
	}

	parts := strings.Split(text, " is ")

	if len(parts) != 2 {
		return errors.New("pattern: invalid input length")
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
		return errors.Wrap(err, "pattern: statement value")
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

// ExecQuestionNumeric parses given text & returns the decimal value of given roman numeral word.
func ExecQuestionNumeric(text string) (string, error) {
	if !patternQuestionNumeric.MatchString(text) {
		return "", errors.New("pattern: input doesn't match rule")
	}

	query := strings.TrimPrefix(text, "how much is ")
	query = strings.TrimSuffix(query, " ?")
	words := strings.Split(query, " ")

	count, err := extractMaterialCount(words)
	if err != nil {
		return "", errors.Wrap(err, "pattern: question numeric")
	}

	return fmt.Sprintf("%s is %d", query, count), nil
}

// ExecQuestionValue parses given text & returns the decimal value of given material & roman numeral word.
func ExecQuestionValue(text string) (string, error) {
	if !patternQuestionValue.MatchString(text) {
		return "", errors.New("pattern: input doesn't match rule")
	}

	query := strings.TrimPrefix(text, "how many Credits is ")
	query = strings.TrimSuffix(query, " ?")
	words := strings.Split(query, " ")

	nums := words[:len(words)-1]
	material := words[len(words)-1]

	count, err := extractMaterialCount(nums)
	if err != nil {
		return "", errors.Wrap(err, "pattern: question value")
	}

	unit, ok := materialUnitMap[material]
	if !ok {
		return "", errors.New("pattern: unknown material")
	}

	value := float64(count) * unit

	return fmt.Sprintf("%s %s is %.0f Credits", strings.Join(nums, " "), material, value), nil
}
