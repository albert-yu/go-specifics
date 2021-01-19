package specifics

import (
	"testing"
)

func TestPartition(t *testing.T) {
	strSlice := []string{
		"dog",
		"cat",
		"bear",
		"octopus",
		"lollipop",
		"dinosaur",
		"leopard",
		"cassowary",
		"lion",
		"mouse",
	}

	// 3 partitions
	expected1 := [][]string{
		strSlice[:4],
		strSlice[4:8],
		strSlice[8:],
	}
	actual1 := StrSlicesPartition(strSlice, 3)
	passed := StrSlicesNestedEqual(actual1, expected1)
	if !passed {
		t.Logf("Expected: %s", expected1)
		t.Logf("Actual: %s", actual1)
		t.Fatal("Partitioning failed")
	}
}
