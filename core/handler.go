package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type execFunc func(text string) (string, error)

var (
	inputFuncs = []execFunc{
		ExecStatementNumeric,
		ExecStatementValue,
		ExecQuestionNumeric,
		ExecQuestionValue,
	}
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your queries:")

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		// Remove multiple spaces.
		input = strings.Join(strings.Fields(input), " ")

		switch input {
		case "":
			continue
		case "exit":
			return
		}

		var output string
		var err error

		for _, fn := range inputFuncs {
			output, err = fn(input)
			if err != nil && err != ErrPatternMismatch {
				fmt.Println("I have no idea what you are talking about")
				break

			} else if err == nil {
				if output != "" {
					fmt.Println(output)
				}
				break
			}
		}

		if err != nil {
			fmt.Println("I have no idea what you are talking about")
		}
	}
}
