package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"flag"
	"os"
)

/* task #9
	1. gather all repositories token has access to
	2. use https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#get-repository-content api endpoint to get directory .github/workflows
	3. retrieve all workflows files
	4. check whether those files contain workflow which is being look for
	4a. [STRETCH/NEXT STEP] compare all workflows with set of workflows
	5. return statistics regarding the usage, which repository uses which workflow
	6. process repositories in goroutines
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
	flag.Parse()

	valid, err := validateControls(controls)
	if !valid {
		fmt.Println("Input is not valid, error: ", err)
		os.Exit(1)
	}
	CONTROLS = &controls
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
		callGitHubAPI(*CONTROLS.PAT)
	}
}

func callGitHubAPI(PAT string) {
	headers := map[string]string{
		"Accept": "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
		"Authorization": fmt.Sprintf("Bearer %s", PAT),
	}
  repos := getRepos(headers)

	files := make([]*GHContent, 0)

	for _, repo := range repos {
		files = append(files, getReposDirectoryContent(repo, ".github/workflows", headers)...)
	}

	rawHeaders := deepCopy(headers)
	rawHeaders["Accept"] = "application/vnd.github.raw+json"
	for idx := range files {
		files[idx].Content = retrieveFileContent(files[idx], rawHeaders)
	}

	filesJsonified, err := json.Marshal(files)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("output.json", filesJsonified, 0777)
	if err != nil {
		fmt.Println(err)
	}
}