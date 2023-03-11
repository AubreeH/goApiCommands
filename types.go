package commands

type CommandHandler = func(args []string)

type Command struct {
	name        string
	handler     CommandHandler
	description string
	aliases     []string
}

type Group struct {
	name     string
	groups   map[string]Group
	commands map[string]Command
	parent   *Group

	compiledCommands map[string]CommandHandler
}
