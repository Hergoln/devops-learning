package std_github

import (
	"fmt"
	"my_module/std_api"
	"my_module/std_json"
	"strings"
)

func GetRepos(headers map[string]string) []*Repository {
	requestURL := "https://api.github.com/user/repos"
	res := std_api.CallAPI(requestURL, headers)
	bodyBytes := std_json.ReadBody(res)
	jsonBody := std_json.UnmarshalArray(bodyBytes)
	return ConvertToRepos(jsonBody)
}

func GetReposDirectoryContent(repo *Repository, directoryPath string, headers map[string]string) []*GHContent {
	directoryUrl := strings.Replace(repo.ContentsUrl, "{+path}", directoryPath, 1)
	res := std_api.CallAPI(directoryUrl, headers)
	if res.StatusCode == 404 {
		fmt.Printf("There is no %s directory in %s repository\n", directoryPath, repo.Name)
		return nil
	}
	bodyBytes := std_json.ReadBody(res)
	jsonBody := std_json.UnmarshalArray(bodyBytes)
	dirContent := ConvertToContent(jsonBody)

	fmt.Printf("Found %d files in %s in %s repository\n", len(dirContent), directoryPath, repo.Name)
	return dirContent
}

func RetrieveFileContent(fileContent *GHContent, headers map[string]string) string {
	res := std_api.CallAPI(fileContent.Download_url, headers)
	if res.StatusCode == 404 {
		fmt.Printf("Error during retrieval of %s", fileContent.Path)
		fmt.Println(res)
		return ""
	}

	return string(std_json.ReadBody(res))
}

func ConvertToContent(jsonBody []map[string]any) []*GHContent {
	workflowsCount := len(jsonBody)
	output := make([]*GHContent, 0, workflowsCount)
	for _, jsonWorkflow := range jsonBody {
		workflow := NewContent(jsonWorkflow)
		if workflow.Type == "file" {
			output = append(output, workflow)
		}
	}

	return output
}

func ConvertToRepos(jsonBody []map[string]any) []*Repository {
	reposCount := len(jsonBody)
	output := make([]*Repository, 0, reposCount)
	for _, jsonRepo := range jsonBody {
		// with type assertion
		output = append(output, NewRepository(jsonRepo))
	}

	return output
}
