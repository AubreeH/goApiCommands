package commands

import (
	"log"
)

func Setup(handler func(group *Group)) Group {
	commands := make(map[string]Command)
	subGroups := make(map[string]Group)

	group := Group{
		name:     "",
		commands: commands,
		groups:   subGroups,
	}

	commands["help"] = Command{
		name:        "Help",
		handler:     group.showHelp,
		description: "Displays a list of available commands",
	}

	handler(&group)
	return group
}

func (group *Group) Group(name string, handler func(group *Group)) {
	_, ok := group.groups[name]
	if !ok {
		commands := make(map[string]Command)
		groups := make(map[string]Group)
		newGroup := Group{
			name:     name,
			commands: commands,
			groups:   groups,
			parent:   group,
		}

		handler(&newGroup)
		group.groups[name] = newGroup
	} else {
		log.Fatalf(`A group with the name "%s" already exists`, name)
	}
}

func (group *Group) Command(handler CommandHandler, name string, description string, aliases ...string) {
	_, ok := group.commands[name]
	if !ok {
		group.commands[name] = Command{
			name:        name,
			description: description,
			handler:     handler,
			aliases:     aliases,
		}
	}
}
