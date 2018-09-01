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

// StringSliceEquals tests the equality between two string slices
func StringSliceEquals(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
