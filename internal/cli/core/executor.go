package core

import (
	"fmt"

	"github.com/octovel/crate/internal/cli/models"
	"github.com/octovel/crate/internal/shared"
)

func Execute(ctx *models.Context) error {
	if !CommandExists(ctx.Input.Name) {
		return fmt.Errorf(
			"Could not find command '%s'.",
			ctx.Input.Name,
		)
	}

	command := ResolveCommand(ctx.Input.Name)

	for _, flag := range command.Flags {
		if flag.Required && !ctx.HasFlag(flag.Name) {
			return fmt.Errorf(
				shared.Ferr("Flag '%s' is required, yet was not found."),
				flag.Name,
			)
		}
	}

	err := command.Execute(ctx)

	return err
}
