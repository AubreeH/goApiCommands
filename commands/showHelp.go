package commands

import (
	"fmt"
	"os"
	"strings"
)

// showHelp prints a list of commands categorised by their groups to the console.
func (group *Group) showHelp(_ []string) {
	fmt.Printf("\nShowing help for %s\n\n%s\n", os.Getenv("PROJECT_NAME"), group.formatHelp(group.name, 0))
}

// Group.formatHelp formats all the Command provided within the targeted group for displaying in showHelp.
func (group *Group) formatHelp(groupPrefix string, indent int) string {
	strIndent := createIndent(indent)

	commandsString := ""

	for _, c := range group.commands {
		if commandsString == "" {
			commandsString += "\n"
		} else {
			commandsString += "\n\n"
		}
		commandsString += c.formatHelp(groupPrefix, strIndent+"\t")
	}
	if commandsString != "" {
		commandsString += "\n"
	}

	groupsString := ""

	for _, g := range group.groups {
		if groupsString == "" {
			if commandsString != "" {
				groupsString += "\n"
			} else {
				groupsString += "\n"
			}
		}

		nextGroupPrefix := groupPrefix
		if groupPrefix != "" && g.name != "" {
			nextGroupPrefix += ":"
		}
		nextGroupPrefix += g.name

		groupsString += strIndent + "\t" + g.formatHelp(nextGroupPrefix, indent+1)
	}

	message := fmt.Sprintf("Group: \"%[2]s\" {%[3]s%[4]s%[1]s}\n", strIndent, group.name, commandsString, groupsString)

	return message
}

// Command.formatHelp formats the targeted command for displaying in GroupStruct.formatHelp and Group.showHelp
func (command *Command) formatHelp(groupPrefix string, strIndent string) string {
	commandNamePart := fmt.Sprintf("%sName: \"%s\"\n", strIndent, command.name)

	exec := groupPrefix
	if exec != "" && command.name != "" {
		exec += ":"
	}
	exec += command.name

	execPart := fmt.Sprintf("%sExec: \"%s\"\n", strIndent, strings.ToLower(exec))

	description := strings.Replace(command.description, "\n", "\n"+strIndent+"             ", -1)
	commandDescriptionPart := fmt.Sprintf(`%sDescription: "%s"`, strIndent, description)
	commandAliasesPart := ""
	if len(command.aliases) > 0 {
		for _, alias := range command.aliases {
			if commandAliasesPart == "" {
				commandAliasesPart += "\n" + strIndent + `Aliases: ["` + alias
			} else {
				commandAliasesPart += `", "` + alias
			}
		}

		if commandAliasesPart != "" {
			commandAliasesPart += `"]`
		}
	}

	return fmt.Sprintf(`%s%s%s%s`, commandNamePart, execPart, commandDescriptionPart, commandAliasesPart)
}

// createIndent creates a string of tabs as long as the provided indent value.
func createIndent(indent int) string {
	return strings.Repeat("\t", indent)
}
