package cli

type command int

const (
	COMMAND_COMMIT command = iota
	COMMAND_COMMIT_ALL
	COMMAND_PUSH
)

type commandFlag string

const (
	COMMAND_COMMIT_FLAG     commandFlag = "c"
	COMMAND_COMMIT_ALL_FLAG             = "ac"
	COMMAND_PUSH_FLAG                   = "p"
)
