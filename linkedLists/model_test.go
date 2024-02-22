package linkedlists

import (
	"testing"

	"github.com/waggertron/data-structures-and-algorithms/common"
)

func TestAdd(t *testing.T) {
	l := Linkedlist{}

	want := 0
	got := l.Length
	if got != want {
		common.ErrorLog(t, got, want)
	}

	l.Add(1)

	want = 1
	got = l.Length
	if got != want {
		common.ErrorLog(t, got, want)
	}

}
