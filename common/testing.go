package common

import "testing"

func ErrorLog(t *testing.T, got, want any) {
	t.Errorf("\nGot: %v\nWant: %v\n", got, want)
}
