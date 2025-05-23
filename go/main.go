package main

import (
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
*/
func main() {
	callGitHubAPI()
}

func callGitHubAPI() {
  requestURL := "https://api.github.com/repos/Hergoln/devops-learning"

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer res.Body.Close()

	fmt.Printf("%d\n", res.StatusCode)

	var body string
	var bodyBytes []byte
	if res.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		body = string(bodyBytes)
	}

	fmt.Printf("Body: %s\n", body)

	var jsonBody map[string]interface{}

	err = json.Unmarshal(bodyBytes, &jsonBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("repo name: %s\n", jsonBody["full_name"].(string))
}