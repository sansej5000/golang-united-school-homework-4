package string_sum

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"fmt"
)

var (
	re                  = regexp.MustCompile(`(?m)([-+]\d+)?([-+]\d+)?([-+]\d+)?`)
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	var one_str, two_str string
	whitespace := strings.ReplaceAll(input, " ", "")
	if len(whitespace) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	for _, match := range re.FindAllStringSubmatch(whitespace, -1) {
		one_str = match[1]
		two_str = match[2]
		if match[2] == "" || match[3] != "" {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		}
	}
	one_dig, err := strconv.Atoi(one_str)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	two_dig, err := strconv.Atoi(two_str)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	output = strconv.Itoa(one_dig + two_dig)
	return output, nil
}

