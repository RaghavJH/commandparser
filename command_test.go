package commandparser

import (
	"testing"
)

var (
	normalCommandStr   = "name arg1 arg2 arg3"
	trailingCommandStr = "   name arg1 arg2 arg3  " //trailing spaces
)

func TestGetSpaceIdxs(t *testing.T) {
	tests := []struct {
		str string
		res []int
	}{
		{normalCommandStr, []int{4, 9, 14}},
		{trailingCommandStr, []int{0, 1, 2, 7, 12, 17, 22, 23}},
	}

	for _, test := range tests {
		idxs := getSpaceIdxs(test.str)
		for i, _ := range idxs {
			if idxs[i] != test.res[i] {
				t.Errorf(
					"Error mismatch: expected %d but got %d\n",
					test.res[i], idxs[i],
				)
			}
		}
	}
}
