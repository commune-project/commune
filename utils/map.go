package utils

// ConcatMaps takes a slice of maps, and put their k/v pairs together.
func ConcatMaps(maps []map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
