package specifics

import (
	"testing"
)

func checkIntArrays(t *testing.T, expected, actual []int, msg string) {
	if !IntSlicesEqual(expected, actual) {
		t.Logf("Arrays are not equal")
		t.Logf("Expected: %v", expected)
		t.Logf("Actual: %v", actual)
		t.Fatal(msg)
	}
}

func TestZipSortedIntArrays(t *testing.T) {
	sorted1 := []int{1, 2, 4, 5}
	sorted2 := []int{1, 3, 7, 10}
	expected := []int{1, 2, 3, 4, 5, 7, 10}
	actual := IntSlicesZipSortedNoDup(sorted1, sorted2)
	checkIntArrays(t, expected, actual, "Zip sorted failed")

	sorted3 := []int{2}
	sorted4 := []int{0}
	expected2 := []int{0, 2}
	actual2 := IntSlicesZipSortedNoDup(sorted3, sorted4)
	checkIntArrays(t, expected2, actual2, "Zip sorted failed")
}

func TestDeleteSorted(t *testing.T) {
	sorted1 := []int{1, 3, 5, 6, 100, 500}
	toDelete1 := 6
	expected1 := []int{1, 3, 5, 100, 500}
	actual1 := IntSlicesDeleteSorted(sorted1, toDelete1)
	checkIntArrays(t, expected1, actual1, "Delete sorted failed")
}
