package main

import "strings"

var implementedCommands = []string{
	"copy",
	"cp",
	"move",
	"mv",
	"remove",
	"rm",
	"mkdir",
	"makedir",
}

func (a *app) cmdAutocomplete(currentText string) (entries []string) {
	if len(currentText) <= 1 {
		return entries
	}
	for _, cmd := range a.commands {
		if strings.HasPrefix(strings.ToLower(cmd), strings.ToLower(currentText)) {
			entries = append(entries, cmd)
		}
	}
	return entries
}
