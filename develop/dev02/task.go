package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const escapeSymbol = 92

func unwrapString(s string) (string, error) {
	runes := []rune(s)
	var builder strings.Builder
	for i := 0; i < len(runes); i++ {
		if runes[i] == escapeSymbol && (i == 0 || runes[i-1] != escapeSymbol) {
			i++
		}
		if i == len(runes)-1 {
			builder.WriteRune(runes[i])
			continue
		}
		if unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+1]) && (i == 0 || runes[i-1] != escapeSymbol) {
			return "", errors.New("incorrect string")
		}
		if unicode.IsDigit(runes[i+1]) {
			num, err := strconv.Atoi(string(runes[i+1]))
			if err != nil {
				return "", err
			}
			for j := 0; j < num; j++ {
				builder.WriteRune(runes[i])
			}
			i++
			continue
		}
		builder.WriteRune(runes[i])
	}

	return builder.String(), nil
}

func main() {
	fmt.Println(unwrapString("a4bc2d5e"))
	fmt.Println(unwrapString("abcd"))
	fmt.Println(unwrapString("45"))
	fmt.Println(unwrapString("qwe/4/5"))
	fmt.Println(unwrapString("qwe/45"))
	fmt.Println(unwrapString("qwe//5"))
}
