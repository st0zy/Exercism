package encode

import (
	"bytes"
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(input); {
		char := string(input[i])
		count := 1
		j := i + 1
		for ; j < len(input) && input[j] == input[i]; j++ {
			count++
		}
		if count != 1 {
			buffer.WriteString(strconv.Itoa(count))
		}
		buffer.WriteString(char)
		i = j
	}

	return buffer.String()
}

func RunLengthDecode(input string) string {
	var result bytes.Buffer
	currentDigit := 1
	prevNumberFound := false
	for _, r := range input {
		if !unicode.IsNumber(r) {
			for _ = range currentDigit {
				result.WriteRune(r)
			}
			prevNumberFound = false
			currentDigit = 1
		} else {
			if prevNumberFound {
				currentDigit = currentDigit*10 + int(r-'0')
			} else {
				currentDigit = int(r - '0')
			}
			prevNumberFound = true
		}
	}

	return result.String()
}
