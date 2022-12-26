package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

const (
	charStart int32 = 97  // a
	charEnd   int32 = 122 // z

	numStart int32 = 48 // 0
	numEnd   int32 = 57 // 9
)

func Unpack(text string) (string, error) {
	if len(text) == 0 {
		return "", nil
	}

	start := int32(text[0])
	if start >= numStart && start <= numEnd {
		return "", ErrInvalidString
	}

	var b strings.Builder
	var tempChar string
	var checkNum bool

	for _, val := range text {
		if checkSymbol(val) {
			return "", ErrInvalidString
		}

		if val >= charStart && val <= charEnd {
			tempChar = string(val) // временно сохраняем элемент текста
		}

		if val >= numStart && val <= numEnd {
			if checkNum {
				return "", ErrInvalidString
			}

			checkNum = true
			count := int(val - '0')

			if count == 0 {
				bText := b.String()
				newText := bText[:len(bText)-len(tempChar)]
				b.Reset()
				b.WriteString(newText)
				continue
			}

			tempChars := strings.Repeat(tempChar, count-1)
			b.WriteString(tempChars)
			tempChar = ""
		} else {
			checkNum = false
			b.WriteString(tempChar)
		}
	}

	return b.String(), nil
}

func checkSymbol(val int32) bool {
	c1 := val >= charStart && val <= charEnd
	c2 := val >= numStart && val <= numEnd

	return !c1 && !c2
}
