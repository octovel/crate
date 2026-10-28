package core

import (
	"os"
	"runtime"
	"time"

	"github.com/octovel/crate/internal/cli/models"
)

func NewContext() *models.Context {
	var state models.State

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
			Build: models.ContextRuntimeBuild{
				Version:   state.Version,
				GitCommit: state.GitCommit,
				GitBranch: state.GitBranch,
				Timestamp: state.BuildTimestamp,
				GoVersion: state.GoVersion,
				Compiler:  state.Compiler,
			},
			WorkingDirectory: cwd,
			ExecutablePath:   exe,
			GOOS:             runtime.GOOS,
			GOARCH:           runtime.GOARCH,
			ProcessID:        os.Getpid(),
			Timestamp:        time.Now(),
		},
		Input: models.ContextInput{
			Name:         "",
			Flags:        []models.ContextFlag{},
			Arguments:    []string{},
			RawArguments: os.Args[1:],
		},
		State: models.State{
			Version:        state.Version,
			GoVersion:      state.GoVersion,
			GitCommit:      state.GitCommit,
			GitBranch:      state.GitBranch,
			BuildTimestamp: state.BuildTimestamp,
			Compiler:       state.Compiler,
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}
