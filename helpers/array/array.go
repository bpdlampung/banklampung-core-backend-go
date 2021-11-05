package array

func InArray(val string, arrays []string) (result string) {
	for x := range arrays {
		if arrays[x] == val {
			result = arrays[x]
		}
	}

	return result
}

func InArrayPointer(val string, arrays []string) (result *string) {
	for x := range arrays {
		if arrays[x] == val {
			result = &arrays[x]
		}
	}

	return result
}
