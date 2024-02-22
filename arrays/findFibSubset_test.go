package arrays

import (
	"reflect"
	"testing"
)

func TestFindFibSubsetEmpty(t *testing.T) {
	want := []int{}
	got := FindFibSubset([]int{})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot: %v\nWant: %v", got, want)
	}
}

func TestFindFibSubsetSimple(t *testing.T) {
	want := []int{3, 2, 1}
	got := FindFibSubset([]int{1, 2, 3})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot: %v\nWant: %v", got, want)
	}
}

func TestFindFibSubset(t *testing.T) {
	want := []int{8, 5, 3}
	got := FindFibSubset([]int{1, 1, 5, 3, 8})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot: %v\nWant: %v", got, want)
	}
}
