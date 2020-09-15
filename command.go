package commandparser

type Command struct {
	cmdStr    string
	spaceIdxs []int
}

func NewCommand(cmdStr string) Command {
	return Command{
		cmdStr:    cmdStr,
		spaceIdxs: make([]int, 0),
	}
}
