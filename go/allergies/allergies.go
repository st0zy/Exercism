package allergies

const (
	eggs = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var _allergens = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(allergies uint) []string {
	initial := 1

	var results []string
	for _, i := range _allergens {
		if allergies&uint(initial) > 0 {
			results = append(results, i)
		}
		initial = initial << 1
	}

	return results

}

func AllergicTo(allergies uint, allergen string) bool {
	for i, candidate := range _allergens {
		if candidate == allergen && (1<<uint(i))&allergies > 0 {
			return true
		}
	}
	return false
}
