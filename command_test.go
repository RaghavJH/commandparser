package commandparser

import (
	"testing"
)

var (
	normalCommandStr   = "name arg1 arg2 arg3"
	trailingCommandStr = "   name arg1 arg2 arg3  " //trailing spaces
	zeroArgCommandStr  = "name"
)

func TestGetArg(t *testing.T) {
	tests := []struct {
		str       string
		name      string
		size      int
		args      []string
		argsError error
	}{
		{
			normalCommandStr,
			"name",
			3,
			[]string{"arg1", "arg2", "arg3"},
			nil,
		},
		{
			trailingCommandStr,
			"name",
			3,
			[]string{"arg1", "arg2", "arg3"},
			nil,
		},
	}

	for _, test := range tests {
		c := NewCommand(test.str)
		if c.Name() != test.name {
			t.Errorf(
				"Error mismatch: expected %s but got %s\n",
				test.name, c.Name(),
			)
		}
		if c.Size() != test.size {
			t.Errorf(
				"Error mismatch: expected %d but got %d\n",
				test.size, c.Size(),
			)
		}

		for i := range test.args {
			arg, _ := c.Arg(i)
			if arg != test.args[i] {
				t.Errorf(
					"Error mismatch: expected %s but got %s\n",
					test.args[i], arg,
				)
			}
		}
	}
}

func TestGetArgErr(t *testing.T) {
	c := NewCommand(zeroArgCommandStr)
	if c.Name() != "name" {
		t.Errorf(
			"Error mismatch: expected %s but got %s\n",
			"name", c.Name(),
		)
	}
	if c.Size() != 0 {
		t.Errorf(
			"Error mismatch: expected %d but got %d\n",
			0, c.Size(),
		)
	}
	_, err := c.Arg(0)
	if err == nil {
		t.Error("Expecting error but received none")
	}
}

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
