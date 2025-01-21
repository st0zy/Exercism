package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

type Set map[string]bool

func New() Set {
	return make(Set, 0)
}

func NewFromSlice(l []string) Set {
	result := New()
	for _, item := range l {
		result.Add(item)
	}
	return result
}

func (s Set) String() string {
	if len(s) == 0 {
		return "{}"
	}
	str := "{"
	for elem := range s {
		str += "\"" + elem + "\", "
	}
	return str[:len(str)-2] + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2) {
		return false
	}

	for item, _ := range s1 {
		if !s2.Has(item) {
			return false
		}
	}
	return true

}

func Disjoint(s1, s2 Set) bool {
	for item, _ := range s1 {
		if s2.Has(item) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	for item, _ := range s1 {
		if !s2.Has(item) {
			return false
		}
	}

	for item, _ := range s2 {
		if !s1.Has(item) {
			return false
		}
	}

	return true

}

func Intersection(s1, s2 Set) Set {
	intersection := New()

	for item, _ := range s1 {
		if s2.Has(item) {
			intersection.Add(item)
		}
	}

	return intersection

}

func Difference(s1, s2 Set) Set {
	difference := New()

	for item, _ := range s1 {
		if !s2.Has(item) {
			difference.Add(item)
		}
	}

	return difference
}

func Union(s1, s2 Set) Set {
	union := New()

	for item, _ := range s1 {
		union.Add(item)
	}

	for item, _ := range s2 {
		union.Add(item)
	}
	return union

}
