package util

import "strings"

// StringSliceContains checks if a slice contains a string
func StringSliceContains(slice []string, search string) bool {
	for _, s := range slice {
		if s == search {
			return true
		}
	}

	return false
}

// StringContainsAny checks if a string contains any of passed parameters
func StringContainsAny(s string, targets ...string) bool {
	for _, t := range targets {
		if strings.Contains(s, t) {
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

// StringSliceRemove returns the slice with the target element removed
func StringSliceRemove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}

	return s
}
