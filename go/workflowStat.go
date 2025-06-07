package main

type WorkflowStat struct {
	Path 				string
	RepoUrl 		string
	References  []*Reference
}

type Reference struct {
	RefRepo string
	Ref string
}

// func toCSV()
// func newWorkflowStat()
