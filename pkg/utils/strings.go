package utils

func IsEmptyString(test string) bool {
	return test == *new(string)
}
