package models

type State struct {
	Version        string
	GoVersion      string
	GitCommit      string
	GitBranch      string
	BuildTimestamp string
	Compiler       string
}
