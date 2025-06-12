package main

import (
	"strings"
	"os"
)

type WorkflowsStatistics struct {
	Stats []*WorkflowStat
	CSVHeaders []string
}

type WorkflowStat struct {
	Path 				string
	RepoUrl 		string
	Usings  []*Usage
}

type Usage struct {
	// whatever comes after "uses:"
	Path string
	// Action, Workflow (?)
	Type string
	Tag string
}

func newWorkflowStat(path string, repoUrl string, usings []*Usage) *WorkflowStat {
	return &WorkflowStat{
		Path: path,
		RepoUrl: repoUrl,
		Usings: usings,
	}
}

func newUsage(path string, type_ string, tag string) *Usage {
	return &Usage{
		Path: path,
		Type: type_,
		Tag: tag,
	}
}

func newStats(stats []*WorkflowStat) *WorkflowsStatistics {
	return &WorkflowsStatistics{
		Stats: stats,
		CSVHeaders: []string{"Workflow Path","Repository URL","Workflow file Path","Type"},
	}
}

func (stat *WorkflowsStatistics) Headers() string {
	return strings.Join(stat.CSVHeaders, ",")
}

// func newWorkflowStat()
func (stat *WorkflowStat) toCSVRows() []string {
	base := make([]string, 0, len(stat.Usings))
	for idx := range stat.Usings {
		row := strings.Join([]string{stat.Path, stat.RepoUrl, stat.Usings[idx].Path, stat.Usings[idx].Type}, ",")
		base = append(base, row)
	}
	return base
}

func (stats *WorkflowsStatistics) StatsToCSV() []string {
	csv := []string{stats.Headers()}
	for idx := range stats.Stats {
		csv = append(csv, stats.Stats[idx].toCSVRows()...)
	}
	return csv
}

func (stats *WorkflowsStatistics) SaveAsCSV() error {
	lines := stats.StatsToCSV()
	// create truncates file
	file, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		check(err)
	}

	return nil
}