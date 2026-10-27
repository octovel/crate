package models

import (
	"io"
	"time"
)

type ContextRuntimeBuild struct {
	Version      string
	GitCommit    string
	GitBranch    string
	Timestamp    string
	GoVersion    string
	Architecture string
	Compiler     string
}

type ContextRuntime struct {
	Build            ContextRuntimeBuild
	WorkingDirectory string
	ExecutablePath   string
	GOOS             string
	GOARCH           string
	ProcessID        int
	Timestamp        time.Time
}

type ContextFlag struct {
	Name    string
	Value   string
	Boolean bool
}

type ContextInput struct {
	Name      string
	Arguments []string
	Flags     []ContextFlag
}

type Context struct {
	Runtime ContextRuntime
	Input   ContextInput
	Stdout  io.Writer
	Stderr  io.Writer
}
