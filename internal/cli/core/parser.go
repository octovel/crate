package core

import (
	"fmt"
	"strings"

	"github.com/octovel/crate/internal/cli/models"
)

func Parse(ctx *models.Context) error {
	args := ctx.Input.RawArguments

	if len(args) < 1 {
		return fmt.Errorf("No command provided.")
	}

	for i := 0; i < len(args); i++ {
		var arg = args[i]

		if !strings.HasPrefix(arg, "-") && ctx.Input.Name == "" {
			ctx.Input.Name = arg
			continue
		}

		if strings.HasPrefix(arg, "-") {
			processed := strings.ToLower(strings.TrimLeft(arg, "-"))

			if processed == "" {
				return fmt.Errorf(
					"Invalid flag format",
				)
			}

			if name, value, found := strings.Cut(processed, "="); found {
				flag := models.ContextFlag{
					Name:  name,
					Value: value,
				}

				if value == "true" || value == "false" {
					flag.Boolean = true
				}

				ctx.Input.Flags = append(ctx.Input.Flags, flag)
				continue
			}

			flag := models.ContextFlag{
				Name: processed,
			}

			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				flag.Value = args[i+1]
				if args[i+1] == "true" || args[i+1] == "false" {
					flag.Boolean = true
				} else {
					flag.Boolean = false
				}
				i++
			} else {
				flag.Value = "true"
				flag.Boolean = true
			}

			ctx.Input.Flags = append(ctx.Input.Flags, flag)
			continue
		} else {
			ctx.Input.Arguments = append(ctx.Input.Arguments, arg)
		}
	}

	return nil
}
