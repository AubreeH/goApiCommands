package commands

import (
	"fmt"
	"strings"
)

// compile runes the compileCommands function on the targeted Group and sets them to the compiledCommands value within the target Group
func (group *Group) compile() error {
	compiledCommands, err := group.compileCommands()
	if err != nil {
		return err
	}
	group.compiledCommands = compiledCommands
	return nil
}

// compileCommands returns a map of CommandHandler indexed by the command needed to run it for every command within the targeted group and all of its subgroups.
func (group *Group) compileCommands() (map[string]CommandHandler, error) {
	compiledCommands := make(map[string]CommandHandler)

	for _, childGroup := range group.groups {
		commands, err := childGroup.compileCommands()
		if err != nil {
			return nil, err
		}
		for n, v := range commands {
			err = compileCommand(compiledCommands, group.name, n, nil, v)
			if err != nil {
				return nil, err
			}
		}
	}

	for n, c := range group.commands {
		err := compileCommand(compiledCommands, group.name, n, c.aliases, c.handler)
		if err != nil {
			return nil, err
		}
	}

	return compiledCommands, nil
}

// compileCommand adds a CommandHandler to the provided map for it's name and every alias provided.
func compileCommand(compiledCommands map[string]CommandHandler, groupName string, commandName string, aliases []string, commandHandler CommandHandler) error {
	name := groupName
	if name != "" && commandName != "" {
		name += ":"
	}
	name += commandName

	handler, ok := compiledCommands[strings.ToLower(name)]
	if ok {
		return fmt.Errorf(`the command name "%s" already in use (Group: "%s", Handler: "%p", Handler In User: "%p") %v`, name+commandName, groupName, commandHandler, handler, compiledCommands)
	}
	compiledCommands[strings.ToLower(name)] = commandHandler

	for _, alias := range aliases {
		name := groupName
		if name != "" && alias != "" {
			name += ":"
		}
		name += alias

		handler, ok := compiledCommands[strings.ToLower(name)]
		if ok {
			return fmt.Errorf(`the alias "%s" is already in use (Command: "%s", Group: "%s", Handler: "%p", Handler In User: "%p") %v`, name+alias, name+commandName, groupName, commandHandler, handler, compiledCommands)
		}
		compiledCommands[strings.ToLower(name)] = commandHandler
	}

	return nil
}
