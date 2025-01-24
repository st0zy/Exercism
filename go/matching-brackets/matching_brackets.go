package brackets

var pairs = map[rune]rune{'{': '}', '[': ']', '(': ')'}

type brackettype int

const (
	openedBracket brackettype = iota
	closedBracket
	notABracket
)

func Bracket(input string) bool {
	var stack []rune

	for _, ch := range input {

		switch getBracketType(ch) {
		case openedBracket:
			stack = append(stack, ch)

		case closedBracket:
			if len(stack) == 0 {
				return false
			}
			lastOpenedBracket := stack[len(stack)-1]
			if ch != pairs[lastOpenedBracket] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0

}

func getBracketType(r rune) brackettype {

	for k, v := range pairs {
		switch r {
		case k:
			return openedBracket
		case v:
			return closedBracket
		}
	}
	return notABracket
}
