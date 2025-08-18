package bubblesort_test

import (
	. "meta/bubblesort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BubbleSort(t *testing.T) {
	arr := []int{1, 2, 3, 5, 4, 7, 9}
	want := []int{1, 2, 3, 4, 5, 7, 9}

	sorted := BubbleSort(arr)

	require.Equal(t, sorted, want)
}

func Test_BubbleSortAllTheWay(t *testing.T) {
	arr := []int{1, 7, 3, 9, 10, 6, 0}
	want := []int{0, 1, 3, 6, 7, 9, 10}

	sorted := BubbleSort(arr)

	require.Equal(t, sorted, want)
}
