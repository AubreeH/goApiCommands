package commands

import (
	"fmt"
	"os"
	"strings"
)

func (group *Group) RunCommand(args ...string) error {
	err := group.compile()
	if err != nil {
		return err
	}

	var cmdArgs []string

	if len(args) > 0 {
		cmdArgs = args
	} else {
		cmdArgs = os.Args
	}

	var commandName string
	var commandArguments []string
	if len(cmdArgs) > 1 {
		commandName = cmdArgs[1]
		if len(cmdArgs) > 2 {
			commandArguments = cmdArgs[1 : len(cmdArgs)-1]
		}
	}

	handler, ok := group.compiledCommands[strings.ToLower(commandName)]
	if ok {
		handler(commandArguments)
	} else {
		fmt.Printf("Unrecognised Command: \"%s\"\nDisplay command info with \"help\" command\n", commandName)
		return fmt.Errorf("Unrecognised Command: \"%s\"\nDisplay command info with \"help\" command\n", commandName)
	}

	return nil
}
