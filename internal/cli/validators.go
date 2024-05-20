package cli

func commandValidator(cmd string) command {
	switch cmd {
	case string(COMMAND_COMMIT_FLAG):
		return COMMAND_COMMIT

	case string(COMMAND_PUSH_FLAG):
		return COMMAND_PUSH

	case string(COMMAND_COMMIT_ALL_FLAG):
		return COMMAND_COMMIT_ALL

	default:
		panic("Invalid command")
	}
}
