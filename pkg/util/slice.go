package util

// StringSliceContains checks if a slice contains a string
func StringSliceContains(slice []string, search string) bool {
	for _, s := range slice {
		if s == search {
			return true
		}
	}

	return false
}
