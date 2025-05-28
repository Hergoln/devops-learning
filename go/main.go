package main

import (
	"os"
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"strings"
)

var (
	REQUEST_ERROR int = 1
	API_CALL_ERROR int = 2
	UNMARSHAL_ERROR int = 3
	BYTE_READ_ERROR int = 4
)

type Repository struct {
	Name string
	Url string
	ContentsUrl string
}

type GHContent struct {
	Name 					string `json:"name"`
	Path 					string `json:"path"`
	Url 					string `json:"url"`
	Git_url 			string `json:"git_url"`
	Download_url 	string `json:"download_url"`
	Type 					string `json:"type"`
	Content				string `json:"content"`
}

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

func callAPI(url string, headers map[string]string) *http.Response {
	fmt.Println("Calling url: ", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(REQUEST_ERROR)
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(API_CALL_ERROR)
	}

	return res
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

	for _, repo := range repos {
		gatherWorkflows(repo, headers)
	}
}

func unmarshalArray(body []byte) []map[string]any {
	var jsonBody []map[string]any

	err := json.Unmarshal(body, &jsonBody)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(UNMARSHAL_ERROR)
	}

	return jsonBody
}

func readBody(response *http.Response) []byte {
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	
	if response.StatusCode == http.StatusOK {	
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(BYTE_READ_ERROR)
		}
		return bytes
	}

	fmt.Println("No body found")
	return nil
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

func newRepository(jsonRepo map[string]any) *Repository {
	repo := Repository{
		Name: jsonRepo["full_name"].(string),
		Url: jsonRepo["url"].(string),
		ContentsUrl: jsonRepo["contents_url"].(string),
	}
	return &repo
}

func gatherWorkflows(repo *Repository, headers map[string]string) []map[string]any {
	workflowsUrl := strings.Replace(repo.ContentsUrl, "{+path}", ".github/workflows", 1)
	res := callAPI(workflowsUrl, headers)
	if res.StatusCode == 404 {
		fmt.Printf("There are no workflows in %s repository\n", repo.Name)
		return nil
	}
	bodyBytes := readBody(res)
	jsonBody := unmarshalArray(bodyBytes)
	workflowsDirectory := extractWorkflows(jsonBody)

	for idx := range workflowsDirectory {
		fmt.Println(workflowsDirectory[idx])
	}

	return nil
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

func newContent(jsonContent map[string]any) *GHContent {
  workflow := GHContent{
		Name: 				jsonContent["name"].(string),
		Path: 				jsonContent["path"].(string),
		Url: 					jsonContent["url"].(string),
		Git_url: 			jsonContent["git_url"].(string),
		Download_url: jsonContent["download_url"].(string),
		Type: 				jsonContent["type"].(string),
	}
	// if workflow.Type == "file" {
	// 	fmt.Println(jsonContent["content"])
	// 	workflow.Content = jsonContent["content"].(string)
	// }
	return &workflow
}