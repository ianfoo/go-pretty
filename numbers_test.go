package pretty

import (
	"strconv"
	"testing"
)

func TestPeriodify(t *testing.T) {
	var tt = []struct {
		input    int
		expected string
	}{
		{input: 1, expected: "1"},
		{input: 10, expected: "10"},
		{input: 100, expected: "100"},
		{input: 1000, expected: "1.000"},
		{input: 10000, expected: "10.000"},
		{input: 100000, expected: "100.000"},
		{input: 1000000, expected: "1.000.000"},
		{input: 1e8, expected: "100.000.000"},
		{input: 1e9, expected: "1.000.000.000"},
		{input: 3478921, expected: "3.478.921"},
	}
	for _, tc := range tt {
		t.Run(strconv.Itoa(tc.input), func(t *testing.T) {
			actual := Periodify(tc.input, '.')
			if actual != tc.expected {
				t.Errorf(`Expected input of %d to yield "%s", but got "%s"`, tc.input, tc.expected, actual)
			}
		})
	}
}

func TestPeriodifyUS(t *testing.T) {
	actual := PeriodifyUS(1000)
	if actual != "1,000" {
		t.Error(`Expected input of 1000 to yield "1,000", but got`, actual)
	}
}
