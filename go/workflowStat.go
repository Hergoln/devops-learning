package main

import (
	"fmt"
	"strings"
)

type WorkflowStat struct {
	Path 				string
	RepoUrl 		string
	Usings  []*Usage
}

type Usage struct {
	// whatever comes after "uses:"
	UsesPath string
	// Action, Workflow (?)
	Type string
}

func newWorkflowStat(path string, repoUrl string) *WorkflowStat {
	return &WorkflowStat{
		Path: path,
		RepoUrl: repoUrl,
	}
}

// func newWorkflowStat()
func (stat *WorkflowStat) toCSVRows() []string {
	base := make([]string, 0, len(stat.Usings))
	for idx := range stat.Usings {
		row := strings.Join([]string{stat.Path, stat.RepoUrl, stat.Usings[idx].UsesPath, stat.Usings[idx].Type}, ",")
		base = append(base, row)
	}
	return base
}