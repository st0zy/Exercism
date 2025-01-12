package wordy

import (
	"fmt"
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	nums, operations := ParseQuestion(question)

	if len(nums) == 0 {
		return 0, false
	}
	if len(nums) == 1 {
		return nums[0], true
	}

	result := 0
	ok := true
	op := 0
	for i, num := range nums {
		if i == 0 {
			result = num
			continue
		}

		result, ok = useOperation(operations[op], result, num)
		op++
		if !ok {
			return 0, false
		}
	}
	return result, ok
}

func useOperation(op string, a, b int) (int, bool) {

	switch op {
	case "plus":
		return a + b, true
	case "minus":
		return a - b, true
	case "multiplied by":
		return a * b, true
	case "divided by":
		return a / b, true

	default:
		return 0, false
	}
}

func ParseQuestion(question string) ([]int, []string) {

	var re *regexp.Regexp
	// re = regexp.MustCompile(`-?[0-9]+ (plus|minus|multiplied by|divided by) -?[0-9]+`)
	// question = re.FindAllString(question, -1)[0]
	// fmt.Println(question)

	re = regexp.MustCompile("(-?[0-9]+)")
	numsAsString := re.FindAllString(question, -1)

	var nums []int
	for _, s := range numsAsString {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic("Failed to parse")
		}
		nums = append(nums, num)
	}
	re = regexp.MustCompile(`-?[0-9]+ (plus|minus|multiplied by|divided by) -?[0-9]+`)
	matches := re.FindStringSubmatch(question)
	operations := matches[1:]
	fmt.Println(matches)
	fmt.Println(operations)
	fmt.Println(nums)
	if len(operations) != len(nums)-1 {
		return nil, nil
	}
	return nums, operations
}
