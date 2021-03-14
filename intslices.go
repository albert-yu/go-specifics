package specifics

// EmptyIntSlice returns an empty, non-nil int slice
func EmptyIntSlice() []int {
	return make([]int, 0)
}

// IntSlicesEqual returns whether two integer arrays are
// equal
func IntSlicesEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, num := range arr1 {
		if num != arr2[i] {
			return false
		}
	}
	return true
}

// IntSlicesZipSortedNoDup zips two already-sorted arrays
// and discards any duplicates
func IntSlicesZipSortedNoDup(intArr1, intArr2 []int) []int {
	var zipped []int
	i := 0
	j := 0
	var lastInt int = -1
	for i < len(intArr1) && j < len(intArr2) {
		currInt1 := intArr1[i]
		currInt2 := intArr2[j]
		if currInt1 < 0 || currInt2 < 0 {
			panic("Got negative number trying to zip sorted")
		}
		for currInt1 == lastInt {
			// this loop should never run
			// more than once, but can in
			// the event of intra-array
			// duplicates
			i++
			currInt1 = intArr1[i]
		}
		for currInt2 == lastInt {
			// ditto for this loop
			j++
			currInt2 = intArr2[j]
		}
		if currInt1 < currInt2 {
			zipped = append(zipped, currInt1)
			lastInt = currInt1
			i++
		} else if currInt2 < currInt1 {
			zipped = append(zipped, currInt2)
			lastInt = currInt2
			j++
		} else {
			zipped = append(zipped, currInt1)
			lastInt = currInt1
			i++
			j++
		}
	}
	// add remaining elements if any
	for ; i < len(intArr1); i++ {
		zipped = append(zipped, intArr1[i])
	}
	for ; j < len(intArr2); j++ {
		zipped = append(zipped, intArr2[j])
	}

	return zipped
}

// IntSlicesBinarySearch returns the index of the found
// element and -1 if not found
func IntSlicesBinarySearch(haystack []int, needle int) int {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return -1
	}

	return low
}

// IntSlicesDeleteSorted deletes the specified value from the int slice
// ONLY if it is found
func IntSlicesDeleteSorted(sortedInt []int, toDelete int) []int {
	foundIndex := IntSlicesBinarySearch(sortedInt, toDelete)
	if foundIndex == -1 {
		return sortedInt
	}
	var newArr []int
	for i, num := range sortedInt {
		if i == foundIndex {
			continue
		}
		newArr = append(newArr, num)
	}

	return newArr
}
