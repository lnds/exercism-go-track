package strain

type Ints []int

func (list Ints) Keep(f func(int) bool) (result Ints) {
	for _, i := range list {
		if f(i) {
			result = append(result, i)
		}
	}
	return result
}

func (list Ints) Discard(f func(int) bool) (result Ints) {
	return list.Keep(func(i int) bool {
		return !f(i)
	})
}

type Strings []string

func (list Strings) Keep(f func(string) bool) (result Strings) {
	for _, i := range list {
		if f(i) {
			result = append(result, i)
		}
	}
	return result
}

type Lists [][]int

func (list Lists) Keep(f func([]int) bool) (result Lists) {
	for _, i := range list {
		if f(i) {
			result = append(result, i)
		}
	}
	return result
}
