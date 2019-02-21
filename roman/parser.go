package roman

import (
	"errors"
)

var (
	ConstantValueMap = map[rune]int64{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

// Parse a given roman number text into decimal. Returns error if a non roman numeral character found.
func Parse(number string) (int64, error) {
	var lastChar rune
	var result int64 = 0

	for _, num := range number {
		value, ok := ConstantValueMap[num]
		if !ok {
			return 0, errors.New("roman: invalid roman character found")
		}

		switch num {
		case 'V', 'X':
			if lastChar == 'I' {
				value -= 2 * ConstantValueMap['I']
			}

		case 'L', 'C':
			if lastChar == 'X' {
				value -= 2 * ConstantValueMap['X']
			}

		case 'D', 'M':
			if lastChar == 'C' {
				value -= 2 * ConstantValueMap['C']
			}
		}

		result += value
		lastChar = num
	}

	return result, nil
}
