package cryptosquare

import (
	"bytes"
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	normalisedString := normalise(pt)

	// fmt.Println("Length of the normalised string is", len(normalisedString))

	r, c := dimensionsForLength(len(normalisedString))
	length := len(normalisedString)

	// fmt.Printf("Dimension required is %d x %d", r, c)

	var encodedMessage bytes.Buffer

	space := " "
	for i := range r * c {
		encodedRow := (i % r)
		encodedColumn := int(i / r)
		encodedIndex := encodedRow*c + encodedColumn
		if i != 0 && encodedRow == 0 {
			encodedMessage.WriteByte([]byte(space)[0])
		}
		if encodedIndex >= length {
			encodedMessage.WriteByte([]byte(space)[0])
		} else {
			encodedMessage.WriteByte(normalisedString[encodedIndex])
		}

	}

	return encodedMessage.String()

}

func normalise(pt string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]")
	pt = reg.ReplaceAllString(pt, "")
	pt = strings.ToLower(pt)
	return pt
}

func dimensionsForLength(len int) (int, int) {
	root, ok := root(len)

	if ok {
		return root, root
	}

	if root*(root+1) >= len {
		return root, root + 1
	}

	return root + 1, root + 1
}

func root(num int) (int, bool) {
	root := int(math.Sqrt(float64(num)))
	if root*root == num {
		return root, true
	}
	return root, false
}
