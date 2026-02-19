package tests

import (
	"slices"
	"testing"
)

type Slice[T comparable] struct{}

func (s Slice[T]) CompareUnorderedSlice(t testing.TB, s1, s2 []T) {
	t.Helper()
	for _, v := range s1 {
		if len(s1) != len(s2) {
			t.Errorf("did not get the right number of elements; want %d, got %d", len(s1), len(s2))
		}
		if !slices.Contains(s2, v) {
			t.Errorf("element '%#v' is not present", v)
		}
	}
}
