package models

import "fmt"

type FlagDefinition struct {
	Name     string
	Value    string
	Boolean  bool
	Required bool
}

type CommandDefinition struct {
	Name        string
	Description string
	Category    string
	Flags       []FlagDefinition
	Action      func(ctx *Context) error
	Validate    func(ctx *Context) error
}

func (c* CommandDefinition) Execute(ctx* Context) error {
	if c.Validate != nil {
		if err := c.Validate(ctx); err != nil {
			return fmt.Errorf(
				"Command '%s' failed to validate: %w",
				c.Name,
				err,
			)
		}
	}

	return c.Action(ctx)
}
