package utils

type StringKeyGetter func(string) string

func DedupeStrings(values []string, keyGetter ...StringKeyGetter) []string {
	getKey := func(input string) string { return input }
	if len(keyGetter) > 0 {
		getKey = keyGetter[0]
	}
	output := []string{}
	seen := map[string]bool{}
	for _, value := range values {
		key := getKey(value)
		if value, ok := seen[key]; ok && value {
			continue
		}
		output = append(output, value)
		seen[key] = true
	}
	return output
}

func IsEmptyString(test string) bool {
	return test == *new(string)
}
