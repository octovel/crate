package core

import "github.com/octovel/crate/internal/cli/models"

type Registry struct {
	Commands map[string]*models.CommandDefinition
	Count    int
}

var CommandsRegistry = Registry{
	Commands: make(map[string]*models.CommandDefinition),
	Count:    0,
}

func Register(cmd *models.CommandDefinition) {
	_, exists := CommandsRegistry.Commands[cmd.Name]
	if !exists {
		CommandsRegistry.Commands[cmd.Name] = cmd
		CommandsRegistry.Count++
	}
}

func CommandExists(name string) bool {
	_, exists := CommandsRegistry.Commands[name]
	return exists
}

func ResolveCommand(name string) *models.CommandDefinition {
	cmd, _ := CommandsRegistry.Commands[name]
	// we don't check if it exists or not here,
	// the executor already does it.
	return cmd
}
