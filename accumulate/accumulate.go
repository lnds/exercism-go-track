package accumulate

func Accumulate(list []string, converter func(string) string) (result []string) {
	for _, s := range list {
		result = append(result, converter(s))
	}
	return
}
