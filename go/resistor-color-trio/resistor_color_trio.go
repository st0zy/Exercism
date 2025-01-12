package resistorcolortrio

import (
	"fmt"
	"math"
)

var bands = map[string]int{

	"black": 0,

	"brown": 1,

	"red": 2,

	"orange": 3,

	"yellow": 4,

	"green": 5,

	"blue": 6,

	"violet": 7,

	"grey": 8,

	"white": 9,
}

type Metric struct {
	name string
	base int
}

var metrics = []Metric{
	{"giga", 1_000_000_000},
	{"mega", 1_000_000},
	{"kilo", 1_000},
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	result, _ := resistance(colors)
	result, unit := convert(result)
	return fmt.Sprintf("%d %sohms", result, unit)
}

func resistance(colors []string) (int, error) {
	resistance := 0

	for _, i := range []int{0, 1} {
		val, ok := bands[colors[i]]
		if !ok {
			return 0, fmt.Errorf("Incorrect color bands specified.")
		}
		resistance = resistance + val*(int(math.Pow10(1-i)))
	}

	multiplier, ok := bands[colors[2]]
	if !ok {
		return 0, fmt.Errorf("Incorrect color bands specified.")
	}

	resistance = resistance * int(math.Pow10(multiplier))
	return resistance, nil
}

func convert(val int) (int, string) {
	if val == 0 {
		return val, ""
	}
	for _, metric := range metrics {
		if val%metric.base == 0 {
			return val / metric.base, metric.name
		}
	}

	return val, ""

}
