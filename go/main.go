package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
)

/* task #9
	1. gather all repositories token has access to
	2. use https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#get-repository-content api endpoint to get directory .github/workflows
	3. retrieve all workflows files
	4. check whether those files contain workflow which is being look for
	4a. [STRETCH/NEXT STEP] compare all workflows with set of workflows
	5. return statistics regarding the usage, which repository uses which workflow
*/
/* task #10
	1. Extract command-wise code to separate locations/files
	2. in main implement parsing of input
	3. validation could be left to each specific command (command pattern?)
	4. implement mechanism which would enable --help flag being used on every command
	5. implement flag to exclude forked repositories
*/
func main() {
	// read from input or read filename from input and then read from file
	PAT := os.Args[1]
	callGitHubAPI(PAT)
}

func callGitHubAPI(PAT string) {
  requestURL := "https://api.github.com/user/repos"
	headers := map[string]string{
		"Accept": "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
		"Authorization": fmt.Sprintf("Bearer %s", PAT),
	}
	res := callAPI(requestURL, headers)
	bodyBytes := readBody(res)
	jsonBody := unmarshalArray(bodyBytes)
	repos := extractRepos(jsonBody)

	workflows := make([]*GHContent, 0)

	for _, repo := range repos {
		workflows = append(workflows, gatherWorkflows(repo, headers)...)
	}

	workflowsJsonified, err := json.Marshal(workflows)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("output.json", workflowsJsonified, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func extractRepos(jsonBody []map[string]any) []*Repository {
	reposCount := len(jsonBody)
	output := make([]*Repository, 0, reposCount)
	for _, jsonRepo := range jsonBody {
		// with type assertion
		output = append(output, newRepository(jsonRepo))
	}

	return output
}

func gatherWorkflows(repo *Repository, headers map[string]string) []*GHContent {
	workflowsUrl := strings.Replace(repo.ContentsUrl, "{+path}", ".github/workflows", 1)
	res := callAPI(workflowsUrl, headers)
	if res.StatusCode == 404 {
		fmt.Printf("There are no workflows in %s repository\n", repo.Name)
		return nil
	}
	bodyBytes := readBody(res)
	jsonBody := unmarshalArray(bodyBytes)
	workflowsDirectory := extractWorkflows(jsonBody)

	return workflowsDirectory
}

func extractWorkflows(jsonBody []map[string]any) []*GHContent {
	workflowsCount := len(jsonBody)
	output := make([]*GHContent, 0, workflowsCount)
	for _, jsonWorkflow := range jsonBody {
		// with type assertion
		output = append(output, newContent(jsonWorkflow))
	}

	return output
}