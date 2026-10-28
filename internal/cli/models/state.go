package models

type State struct {
	Version        string `json:"version"`
	GoVersion      string `json:"go_version"`
	GitCommit      string `json:"git_commit"`
	GitBranch      string `json:"git_branch"`
	BuildTimestamp string `json:"build_timestamp"`
	Compiler       string `json:"compiler"`
}
