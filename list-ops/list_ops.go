package listops

type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

type IntList []int

func (list IntList) Foldr(fn binFunc, acum int) int {
	for i := len(list) - 1; i >= 0; i-- {
		acum = fn(list[i], acum)
	}
	return acum
}

func (list IntList) Foldl(fn binFunc, acum int) int {
	for _, e := range list {
		acum = fn(acum, e)
	}
	return acum
}

func (list IntList) Filter(fn predFunc) IntList {
	result := IntList{}
	for _, e := range list {
		if fn(e) {
			result = append(result, e)
		}
	}
	return result
}

func (list IntList) Length() int {
	return len(list)
}

func (list IntList) Map(fn unaryFunc) IntList {
	result := IntList{}
	for _, e := range list {
		result = append(result, fn(e))
	}
	return result
}

func (list IntList) Reverse() IntList {
	result := IntList{}
	for _, e := range list {
		result = append(IntList{e}, result...)
	}
	return result

}

func (list IntList) Append(other IntList) IntList {
	return append(list, other...)
}

func (list IntList) Concat(lists []IntList) IntList {
	result := IntList{}
	result = append(result, list...)
	for _, l := range lists {
		result = append(result, l...)
	}
	return result
}
