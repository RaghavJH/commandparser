package commandparser

import (
	"errors"
	"strings"
)

var (
	ErrArgOutOfBounds = errors.New("Argument index out of bounds")
)
//Hello world
type Command struct {
	cmdStr    string
	spaceIdxs []int
}

func NewCommand(cmdStr string) Command {
	cmdStr = strings.Trim(cmdStr, " ")
	return Command{
		cmdStr:    cmdStr,
		spaceIdxs: getSpaceIdxs(cmdStr),
	}
}

//Returns the name of the command e.g.
//command is "levelup 1", it returns "levelup"
func (c *Command) Name() string {
	//If there is no space, the command is
	//argless i.e. just the name on its own
	if len(c.spaceIdxs) == 0 {
		return c.cmdStr
	}

	return c.cmdStr[:c.spaceIdxs[0]]
}

//Returns the arg at index i (starting at 0)
//Error if out of bounds
func (c *Command) Arg(i int) (string, error) {
	if c.Size() == 0 || i < 0 || i >= c.Size() {
		return "", ErrArgOutOfBounds
	}

	//Edge case for last argument
	if i == c.Size()-1 {
		return c.cmdStr[c.spaceIdxs[i]+1:], nil
	}
	return c.cmdStr[c.spaceIdxs[i]+1 : c.spaceIdxs[i+1]], nil
}

//Returns the number of args (not including name)
func (c *Command) Size() int {
	return len(c.spaceIdxs)
}

//Returns the indexes of each space in
//a string. Use this as opposed to string.Split
//as a means to extract args to reduce
//GC pressure significantly (40% less memory)
//http://prntscr.com/uhc4m7
func getSpaceIdxs(s string) []int {
	spaces := make([]int, 0, len(s))
	space := strings.Index(s, " ")
	i := space
	for space >= 0 {
		spaces = append(spaces, i)
		space = strings.Index(s[i+1:], " ")
		i += space + 1
	}
	return spaces
}
