package models

import "fmt"

type Flag struct {
	Name        string
	Description string
	Values      []string
	Required    bool
	Type        string // "bool" or "str"
}

type Command struct {
	Name        string
	Description string
	Group       string
	Flags       []Flag
	Action      func(ctx *Context) error
	Validate    func(ctx *Context) error
}

func (cmd *Command) Execute(ctx *Context) error {
	if cmd == nil {
		return fmt.Errorf("Tried executing a 'nil' command.")
	}

	if cmd.Action == nil {
		return fmt.Errorf("Command '%s' is missing a 'Action()' function.", cmd.Name)
	}

	if cmd.Validate != nil {
		if err := cmd.Validate(ctx); err != nil {
			return fmt.Errorf("Command '%s' validation failed: %w", cmd.Name, err)
		}
	}

	return cmd.Action(ctx)
}
