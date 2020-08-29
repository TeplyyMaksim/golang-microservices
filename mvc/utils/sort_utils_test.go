package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// Worst case because algorithm should  replace every element
func getWorstCaseElements(n int, initialElementValue int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = initialElementValue - i
	}

	return result
}

func TestBubbleSortWorthCase(t *testing.T) {
	elements := getWorstCaseElements(5, 9)
	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

// Best case because no replace required
func TestBubbleSortBestCase(t *testing.T) {
	elements := []int{5, 6, 7, 8, 9}
	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

// Nil check case
func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)
}

func BenchmarkBubbleSort1000Elements(b *testing.B) {
	elements := getWorstCaseElements(1000, 1000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort1000 (b *testing.B) {
	elements := getWorstCaseElements(1000, 1000)

	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}