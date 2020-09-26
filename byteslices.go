package specifics

import (
	"math"
)

// ByteSlicesEqual determines if two
// byte slices are equal
func ByteSlicesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// ByteSlicesPartition partitions the slice of byte slices into `k` subslices
func ByteSlicesPartition(slice [][]byte, k int) [][][]byte {
	var result [][][]byte
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
