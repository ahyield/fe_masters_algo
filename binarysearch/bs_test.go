package binarysearch_test

import (
	"testing"

	. "meta/binarysearch"

	"github.com/stretchr/testify/require"
)

func TestBS(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 3, 5, 9, 10}

	for _, v := range want {
		_, err := BinarySearch(arr, v)
		require.NoError(t, err)
	}
}

func TestBS_Fail(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{29, 33, 15, 95, 11}

	for _, v := range want {
		_, err := BinarySearch(arr, v)
		require.Error(t, err)
		require.ErrorContains(t, err, "Number doesn't exist")
	}
}
