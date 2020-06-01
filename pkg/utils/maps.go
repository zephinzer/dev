package utils

// GetNKeyValuePairsStringMap returns the number of key-value mappings stored
// in the provided `kvd` (key-value dictionary)
func GetNKeyValuePairsStringMap(kvd map[string]string) int {
	count := 0
	for range kvd {
		count++
	}
	return count
}
