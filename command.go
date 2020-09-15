package commandparser

import "strings"

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
