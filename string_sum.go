package string_sum
//packege strconv

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

//используйте эти ошибки соответствующим образом, обернув их функцией fmt.Errorf
var (
	re = regexp.MustCompile(`(?m)([-+]{0,})\s{0,}(\d{1,})\s{0,}([-+]{1,})\s{0,}(\d{1,})`)
	// Используйте, когда ввод пуст, и ввод считается пустым, если строка содержит только пробелы
	errorEmptyInput = errors.New("input is empty")
	// Используйте, когда выражение имеет количество операндов, не равное двум
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Реализуйте функцию, которая вычисляет сумму двух целых чисел, записанных в виде строки.
// Например, имея входную строку «3+5», он должен возвращать выходную строку «8» и нулевую ошибку.
// Рассмотрим случаи, когда операнды отрицательные ("-3+5" или "-3-5") и когда входная строка содержит пробелы ("3 + 5").
//
//Для случаев, когда входное выражение неверно (содержит символы, не являющиеся числами, +, - или пробел)
// функция должна возвращать пустую строку и соответствующую ошибку из пакета strconv, завернутую в вашу собственную ошибку
// с функцией fmt.Errorf
//
// Используйте ошибки, определенные выше, как описано, снова завернув в fmt.Errorf

func StringSum(input string) (output string, err error) {

	var one, two, three, four string

	whitespace := strings.ReplaceAll(input, " ", "")

	fmt.Println(len(whitespace))

	if len(whitespace) == 0 {
		fmt.Println("input is empty")
	}

	for _, match := range re.FindAllStringSubmatch(input, -1) {
		one = match[1]
		two = match[2]
		three = match[3]
		four = match[4]
	}


	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)

	fmt.Println(strings.HasPrefix(input, "-"))
	return "", nil
}

func main() {
	itog := StringSum("-3+5")
	fmt.Println(itog)
}
