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

// Contains detects if e is in slice s
func Contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
