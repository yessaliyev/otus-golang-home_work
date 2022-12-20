package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

const (
	charStart int32 = 97  // a
	charEnd   int32 = 122 // z

	numStart int32 = 48 // 0
	numEnd   int32 = 57 // 9
)

func validate(chars string) bool {
	if len(chars) == 0 {
		return true
	}

	start := int32(chars[0])

	if start >= numStart && start <= numEnd {
		return false
	}

	var checkNum bool

	for _, char := range chars {
		if char >= numStart && char <= numEnd {
			if checkNum == true {
				return false
			}
			checkNum = true
			continue
		}

		if char >= charStart && char <= charEnd {
			checkNum = false
			continue
		}
	}

	return true
}

//для Unpack2
func counterForSlice(chars string) int {
	if len(chars) == 0 {
		return 0
	}

	var counter int

	for _, char := range chars {
		if char >= numStart && char <= numEnd {
			counter += int(char-'0') - 1
		}

		if char >= charStart && char <= charEnd {
			counter++
		}
	}

	return counter
}

func Unpack(chars string) (string, error) {
	if !validate(chars) {
		return "", ErrInvalidString
	}

	var tempChar string
	var charSlice []string

	for _, char := range chars {

		if char >= charStart && char <= charEnd {
			tempChar = string(char)
			charSlice = append(charSlice, string(char))
			continue
		}

		if char == numStart {
			charSlice = charSlice[:len(charSlice)-1]
		}

		if char > numStart && char <= numEnd {
			counter, _ := strconv.Atoi(string(char))
			charSlice = append(charSlice, strings.Repeat(tempChar, counter-1))
		}
	}

	return strings.Join(charSlice, ""), nil
}

func Unpack2(chars string) (string, error) {
	if !validate(chars) {
		return "", ErrInvalidString
	}

	var tempChar string
	charSlice := make([]string, counterForSlice(chars))

	for _, char := range chars {

		if char >= charStart && char <= charEnd {
			tempChar = string(char)
			charSlice = append(charSlice, string(char))
			continue
		}

		if char == numStart {
			charSlice = charSlice[:len(charSlice)-1]
		}

		if char > numStart && char <= numEnd {
			counter, _ := strconv.Atoi(string(char))
			charSlice = append(charSlice, strings.Repeat(tempChar, counter-1))
		}
	}

	return strings.Join(charSlice, ""), nil
}

func Unpack3(text string) (string, error) {
	if !validate(text) {
		return "", ErrInvalidString
	}

	var tempChar string

	newText := ""

	for _, val := range text {
		if val >= charStart && val <= charEnd {
			tempChar = string(val) // временно сохраняем элемент текста
		}

		if val >= numStart && val <= numEnd {

			if tempChar == "" {
				fmt.Println("error")
				break
			}

			count := int(val - '0')

			if count == 0 {
				newText = newText[:len(newText)-len(tempChar)]
				continue
			}

			tempChars := ""

			for i := 1; i < count; i++ {
				tempChars += tempChar
			}

			newText += tempChars
			tempChar = ""
		} else {
			newText += tempChar
		}
	}

	return newText, nil
}

func Unpack4(text string) (string, error) {
	if !validate(text) {
		return "", ErrInvalidString
	}

	var b strings.Builder
	var tempChar string

	newText := ""

	for _, val := range text {
		if val >= charStart && val <= charEnd {
			tempChar = string(val) // временно сохраняем элемент текста
		}

		if val >= numStart && val <= numEnd {

			if tempChar == "" {
				fmt.Println("error")
				break
			}

			count := int(val - '0')

			if count == 0 {
				bText := b.String()
				newText = bText[:len(bText)-len(tempChar)]
				b.Reset()
				b.WriteString(newText)
				continue
			}

			tempChars := ""

			tempChars = strings.Repeat(tempChar, count-1)
			b.WriteString(tempChars)
			tempChar = ""
		} else {
			b.WriteString(tempChar)
		}
	}

	return b.String(), nil
}
