package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		exp := strings.Fields(s.Text())
		if len(exp) != 3 {
			panic("не правильное выражение")
		}
		num1 := exp[0]
		num2 := exp[2]
		oper := exp[1]

		num1Arab, err1 := strconv.Atoi(num1)
		num2Arab, err2 := strconv.Atoi(num2)

		if err1 == nil && err2 == nil {
			fmt.Println(computationt(num1Arab, num2Arab, oper))
			continue
		}

		val1 := isValidRoman(num1)
		val2 := isValidRoman(num2)

		if val1 && val2 {
			num1Arab = Decode(num1)
			num2Arab = Decode(num2)
			result := computationt(num1Arab, num2Arab, oper)
			if result < 1 {
				panic("not minus roman")
			}

			fmt.Println(intToRoman(result))
			continue
		}

		panic("не правильное выражение")
	}
}

func computationt(num1Arab, num2Arab int, oper string) (result int) {

	switch oper {
	case "-":
		result = num1Arab - num2Arab
	case "+":
		result = num1Arab + num2Arab
	case "/":
		result = num1Arab / num2Arab
	case "*":
		result = num1Arab * num2Arab
	default:
		panic("not oper")
	}

	return
}

func Decode(roman string) int {
	var sum int
	var Roman = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for k, v := range roman {
		if k < len(roman)-1 && Roman[byte(roman[k+1])] > Roman[byte(roman[k])] {
			sum -= Roman[byte(v)]
		} else {
			sum += Roman[byte(v)]
		}
	}
	return sum
}

func isValidRoman(s string) bool {
	// Регулярное выражение для проверки римского числа
	re := regexp.MustCompile(`^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	return re.MatchString(s)
}

func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var roman strings.Builder

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman.WriteString(syb[i])
		}
	}

	return roman.String()
}
