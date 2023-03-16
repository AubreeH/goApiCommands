package commands

// CommandHandler is a function that takes an array of strings
type CommandHandler = func(args []string)

// Command is the parsed value created when running Group.Command
type Command struct {
	name        string
	handler     CommandHandler
	description string
	aliases     []string
}

// Group is the parsed value returned when running Setup and created when running Group.Group
type Group struct {
	name     string
	groups   map[string]Group
	commands map[string]Command
	parent   *Group

	compiledCommands map[string]CommandHandler
}
