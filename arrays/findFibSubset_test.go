package arrays

import (
	"reflect"
	"testing"
)

func TestFindFibSubset(t *testing.T) {
	want := []int{}
	got := FindFibSubset([]int{})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot: %v\nWant: %v", got, want)
	}
}
