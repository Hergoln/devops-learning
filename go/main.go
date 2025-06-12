package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"unicode"
)

/* task #9
	5. process repositories in goroutines -> requires to process repo by repo
*/
/* task #10
Cobra framework - https://github.com/spf13/cobra
	1. Extract command-wise code to separate locations/files
	2. in main implement parsing of input
	3. validation could be left to each specific command (command pattern?)
	4. implement mechanism which would enable --help flag being used on every command
	5. implement flag to exclude forked repositories
*/

var (
	CONTROLS *Control
)

func init() {
	controls := Control{}
	controls.PAT = flag.String("pat", "", "GitHub PAT token")
	controls.CMD = flag.String("cmd", "", "Which command is being run")
	controls.WF_REPO = flag.String("workflows_repo", "", "repository ({OWNER}/{REPO NAME}) that workflows are being checked against")
	flag.Parse()

	valid, err := validateControls(controls)
	if !valid {
		fmt.Println("Input is not valid, error: ", err)
		os.Exit(1)
	}
	CONTROLS = &controls
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func deepCopy(copied map[string]string) map[string]string {
	copy := map[string]string{}
	for key, value := range copied {
		copy[key] = value
	}
	return copy
}

func main() {
	// read from input or read filename from input and then read from file
	if *CONTROLS.CMD == STATS_CMD {
		gatherWorkflowsStats(CONTROLS)
	}
}

func gatherWorkflowsStats(CONTROLS *Control) {
	headers := map[string]string{
		"Accept": "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
		"Authorization": fmt.Sprintf("Bearer %s", *CONTROLS.PAT),
	}
  repos := getRepos(headers)
	rawHeaders := deepCopy(headers)
	rawHeaders["Accept"] = "application/vnd.github.raw+json"
	stats := make([]*WorkflowStat, 0)

	for _, repo := range repos {
		files := getReposDirectoryContent(repo, ".github/workflows", headers)
		for fileIdx := range files {
			files[fileIdx].Content = retrieveFileContent(files[fileIdx], rawHeaders)
			uses := extractUses(files[fileIdx].Content)
			stats = append(stats, newWorkflowStat(files[fileIdx].Path, repo.Url, uses))
		}
	}

	stat := newStats(stats)

	err := stat.SaveAsCSV()
	check(err)
}

func extractUses(content string) []*Usage {
	uses := make([]*Usage, 0)
	for l := range strings.Lines(content) {
		usesIdx := strings.Index(l, "uses:")

		if usesIdx > -1 {
			line := strings.TrimSpace(l[usesIdx+len("uses:"):])
			tagIdx := strings.Index(line, "@")
			path := line[:tagIdx]
			line = line[tagIdx+1:]
			commentIdx := strings.IndexFunc(line, unicode.IsSpace)
			if commentIdx > -1 {
				uses = append(uses, newUsage(path, "", line[:commentIdx]))
			} else {
				uses = append(uses, newUsage(path, "", line))
			}
		}
	}
	return uses
}