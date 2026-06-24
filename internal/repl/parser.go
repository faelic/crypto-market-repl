package repl

import "strings"

func parseInput(input string) (string, []string) {

	fields := strings.Fields(input)

	if len(fields) == 0 {
		return "", []string{}
	}
	command := fields[0]
	args := fields[1:]

	return command, args
}
