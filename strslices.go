package specifics

import "math"

// EmptyStrSlice returns an empty string slice
func EmptyStrSlice() []string {
	return make([]string, 0)
}

// StrSliceContains checks if the string slice
// contains the given string
func StrSliceContains(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}

// StrSlicesMap is a `map` for string slices
func StrSlicesMap(strarr []string, lambda func(s string) string) []string {
	result := make([]string, len(strarr))
	for i, elem := range strarr {
		result[i] = lambda(elem)
	}
	return result
}

// StrSlicesEqual tests if two string slices
// are equivalent
func StrSlicesEqual(a, b []string) bool {
	if a == nil != (b == nil) {
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

// StrSlicesNestedEqual tests if two slices of
// string slices are equal
func StrSlicesNestedEqual(a, b [][]string) bool {
	if a == nil != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !StrSlicesEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

// StrSlicesPartition partitions the slice into `k` subslices
func StrSlicesPartition(slice []string, k int) [][]string {
	var result [][]string
	n := len(slice)
	sizes := make([]int, k)
	if n%k == 0 {
		subsliceSize := n / k
		for i := 0; i < k; i++ {
			sizes[i] = subsliceSize
		}
	} else {
		// do ceiling division
		chunkSizeFl := float64(n) / float64(k)
		chunkSize := int(math.Ceil(chunkSizeFl))
		for i := 0; i < k-1; i++ {
			sizes[i] = chunkSize
		}

		// remainder is size of last chunk
		remainder := n % chunkSize
		sizes[k-1] = remainder
	}

	// make the subslices
	index := 0
	for _, sz := range sizes {
		subslice := slice[index : index+sz]
		result = append(result, subslice)
		index += sz
	}
	return result
}

// StrPtrSlicesToStrSlices converts a slice of string
// pointers to strings
func StrPtrSlicesToStrSlices(input []*string) []string {
	output := make([]string, len(input))
	for i, ptr := range input {
		if ptr == nil {
			output[i] = ""
		} else {
			output[i] = *ptr
		}
	}
	return output
}

// StrSliceRemoveString removes the all occurrences, if any, of the target string
func StrSliceRemoveString(slice []string, toRemove string) []string {
	var result []string
	for _, s := range slice {
		if s == toRemove {
			continue
		}
		result = append(result, s)
	}
	return result
}
