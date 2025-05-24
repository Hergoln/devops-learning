package main

import (
	"os"
	"fmt"
	"net/http"
	"time"
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

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", PAT))

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bodyBytes := readBody(res)
	jsonBody := unmarshalReposRes(bodyBytes)
	names := extractRepos(jsonBody)

	fmt.Println(names)
}

func unmarshalReposRes(body []byte) []map[string]any {
	// var jsonBody map[string]any
	var jsonBody []map[string]any

	err := json.Unmarshal(body, &jsonBody)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// fmt.Printf("repo name: %s\n", jsonBody["full_name"].(string))
	return jsonBody
}

func readBody(response *http.Response) []byte {
	defer response.Body.Close()
	fmt.Printf("%d\n", response.StatusCode)

	if response.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return bytes
	}

	return nil
}

func extractRepos(jsonBody []map[string]any) []string {
	reposCount := len(jsonBody)
	output := make([]string, 0, reposCount)
	for idx := range jsonBody {
		// with type assertion
		output = append(output, jsonBody[idx]["full_name"].(string))
	}

	return output
}