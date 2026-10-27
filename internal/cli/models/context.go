package models

import (
	"io"
	"time"
)

type Context struct {
	Command ContextCommand
	Runtime ContextRuntime
	Stdout  io.Writer
	Stderr  io.Writer
}

type ContextCommand struct {
	Name      string
	Arguments []string
	Flags     map[string]ContextFlag
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
	Name  string
	Value string
	Type  string // "bool" or "str"
}

type ContextRuntimeBuild struct {
	Version      string
	GitCommit    string
	GitBranch    string
	Timestamp    string
	GoVersion    string
	Architecture string
	Compiler     string
}
