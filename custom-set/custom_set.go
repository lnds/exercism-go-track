package stringset

import (
	"fmt"
	"strings"
)

type void struct{}
type Set map[string]void

var member void

func New() Set {
	return make(Set)
}

func NewFromSlice(elements []string) Set {
	result := make(Set)
	for _, e := range elements {
		result[e] = member
	}
	return result
}

func (s Set) String() string {
	keys := make([]string, 0, len(s))
	for key := range s {
		keys = append(keys, fmt.Sprintf("\"%s\"", key))
	}
	return fmt.Sprintf("{%s}", strings.Join(keys, ", "))
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = member
}

func Subset(s1, s2 Set) bool {
	return len(Intersection(s1, s2)) == len(s1)
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		if _, ok := s2[k]; ok {
			result[k] = member
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			result[k] = member
		}
	}
	return result
}

func Union(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		result[k] = member
	}
	for k := range s2 {
		result[k] = member
	}
	return result
}
