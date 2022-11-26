package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFlightPath(t *testing.T) {
	testCases := []struct {
		start, end string
		flights    []flight
		expect     bool
	}{
		{
			"A",
			"B",
			[]flight{
				{
					"A",
					"B",
					1,
					2,
				},
			},
			true,
		},
		{
			"A",
			"C",
			[]flight{
				{
					"A",
					"B",
					1,
					2,
				},
			},
			false,
		},
		{
			"A",
			"C",
			[]flight{
				{
					"A",
					"B",
					-1,
					2,
				},
			},
			false,
		},
	}

	for _, testCase := range testCases {
		got := GetFlightPath(testCase.start, testCase.end, testCase.flights)

		assert.Equal(t, testCase.expect, got)
	}
}
