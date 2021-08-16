package flatten

func Flatten(input interface{}) []interface{} {
	result := []interface{}{}
	if list, ok := input.([]interface{}); ok {
		for _, e := range list {
			switch elem := e.(type) {
			case []interface{}:
				result = append(result, Flatten(elem)...)
			case interface{}:
				result = append(result, elem)
			}
		}
	}
	return result
}
