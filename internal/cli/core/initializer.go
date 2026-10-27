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
		Runtime: models.ContextRuntime{
			WorkingDirectory: cwd,
			ExecutablePath:   exe,
			GOOS:             runtime.GOOS,
			GOARCH:           runtime.GOARCH,
			ProcessID:        os.Getpid(),
			Timestamp:        time.Now(),
		},
		Input: models.ContextInput{
			Name:      "",
			Arguments: []string{},
			Flags:     []models.ContextFlag{},
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}
