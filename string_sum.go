package string_sum

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"fmt"
)

var (
	re = regexp.MustCompile(`(?m)([-+]{0,})([\d\w]{1,})([-+]{1,})([\d\w]{1,})`)
	errorEmptyInput = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	var one_str, two_str, three_str, four_str string
	var sum int
	whitespace := strings.ReplaceAll(input, " ", "")
	if len(whitespace) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	} else {
		for _, match := range re.FindAllStringSubmatch(whitespace, -1) {
			one_str = match[1]
			two_str = match[2]
			three_str = match[3]
			four_str = match[4]
		}
		if one_str != "-" && one_str != "+" && one_str != "" || three_str != "-" && three_str != "+" {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		} else {
			two_dig, err := strconv.Atoi(one_str + two_str)
			if err != nil {
				return "", fmt.Errorf("%w", err)
			}
			four_dig, err := strconv.Atoi(three_str + four_str)
			if err != nil {
				return "", fmt.Errorf("%w", err)
			} else {
				sum = two_dig + four_dig
				output = strconv.Itoa(sum)
				return output, nil
			}
		}
	}
}

