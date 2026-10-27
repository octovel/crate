package core

import (
	"os"
	"runtime"
	"time"

	"github.com/octovel/crate/internal/cli/models"
)

func NewContext() *models.Context {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "unknown"
	}

	exe, err := os.Executable()
	if err != nil {
		exe = "unknown"
	}

	return &models.Context{
		Command: models.ContextCommand{
			Name:      "",
			Arguments: []string{},
			Flags:     make(map[string]models.ContextFlag),
		},
		Runtime: models.ContextRuntime{
			WorkingDirectory: cwd,
			ExecutablePath:   exe,
			GOOS:             runtime.GOOS,
			GOARCH:           runtime.GOARCH,
			ProcessID:        os.Getpid(),
			Timestamp:        time.Now(),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}
