package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const EscapeRune = 92

func Unpack(s string) (string, error) {
	var result strings.Builder
	var escape, escapeExist bool

	for i, value := range s {
		switch {
		case escape && unicode.IsLetter(value):
			return "", ErrInvalidString

		case escape:
			result.WriteRune(value)
			escape = false

		case !escape && value == EscapeRune:
			escape, escapeExist = true, true

		case unicode.IsDigit(value) && len(result.String()) > 0:
			number, err := strconv.Atoi(string(value))
			if !escapeExist && unicode.IsDigit(rune(s[i-1])) && err == nil {
				return "", ErrInvalidString
			}

			if number == 0 {
				resultCopy := result.String()
				result.Reset()
				result.WriteString(resultCopy[:len(resultCopy)-1])
			} else {
				resultCopy := result.String()
				result.WriteString(strings.Repeat(resultCopy[len(resultCopy)-1:], number-1))
			}

		case unicode.IsLetter(value):
			result.WriteRune(value)

		default:
			return "", ErrInvalidString
		}
	}

	return result.String(), nil
}
