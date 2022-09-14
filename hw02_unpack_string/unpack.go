package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func checkFirstIsDigit(packedString string) bool {
	// true если первый элемент число
	return packedString != "" && unicode.IsDigit(rune(packedString[0]))
}

func checkCurrentIsDigitAndPreviousNotBackslash(item rune, index int, sr []rune) bool {
	// true если текущий элемент число, предыдущий элемент число и пред предыдущий
	return unicode.IsDigit(item) && unicode.IsDigit(sr[index-1]) && sr[index-2] != '\\'
}

func checkBackslash(item rune, backslash bool) bool {
	// true если элемент равен \\
	return item == '\\' && !backslash
}

func checkBackslashAndIsLetter(backslash bool, item rune) bool {
	// true если предущий элемент был // и текущий элемент буква
	return backslash && unicode.IsLetter(item)
}

func Unpack(s string) (string, error) {
	sr := []rune(s)
	builder := strings.Builder{}
	var n int
	var backslash bool

	if checkFirstIsDigit(s) {
		return "", ErrInvalidString
	}

	for index, item := range sr {
		if checkCurrentIsDigitAndPreviousNotBackslash(item, index, sr) {
			return "", ErrInvalidString
		}
		if checkBackslash(item, backslash) {
			backslash = true
			continue
		}
		if checkBackslashAndIsLetter(backslash, item) {
			return "", ErrInvalidString
		}
		if backslash {
			builder.WriteString(string(item))
			backslash = false
			continue
		}
		if unicode.IsDigit(item) {
			n = int(item - '0')
			if n == 0 {
				unpackedString := builder.String()
				builder.Reset()
				builder.WriteString(unpackedString[:len(unpackedString)-1])
				continue
			}
			builder.WriteString(strings.Repeat(string(sr[index-1]), n-1))
			continue
		}
		builder.WriteString(string(item))
	}

	return builder.String(), nil
}
