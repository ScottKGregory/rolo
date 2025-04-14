package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/scottkgregory/rolo/internal/assert"
)

func Test(tt *testing.T) {
	testCases := []struct {
		input       []string
		expectedMin int
		expectedMax int
		expectedErr error
	}{
		{
			input:       []string{"1d4"},
			expectedMin: 0,
			expectedMax: 4,
		},
		{
			input:       []string{"2d20"},
			expectedMin: 0,
			expectedMax: 40,
		},
		{
			input:       []string{"(1d5+4)+1"},
			expectedMin: 1,
			expectedMax: 10,
		},
	}
	for _, tc := range testCases {
		tt.Run(strings.Join(tc.input, " "), func(t *testing.T) {
			for range 500 {
				actual, err := run(tc.input)
				assert.ErrorEqual(t, tc.expectedErr, err)

				ai, err := strconv.Atoi(actual)
				assert.Nil(t, err)

				assert.Equal(t, true, ai >= tc.expectedMin && ai <= tc.expectedMax)
			}
		})
	}
}
