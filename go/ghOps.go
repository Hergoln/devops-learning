package main

import (
	"fmt"
	"strings"
)

func getRepos(headers map[string]string) []*Repository {
	requestURL := "https://api.github.com/user/repos"
	res := callAPI(requestURL, headers)
	bodyBytes := readBody(res)
	jsonBody := unmarshalArray(bodyBytes)
	return convertToRepos(jsonBody)
}

func getReposDirectoryContent(repo *Repository, directoryPath string, headers map[string]string) []*GHContent {
	directoryUrl := strings.Replace(repo.ContentsUrl, "{+path}", directoryPath, 1)
	res := callAPI(directoryUrl, headers)
	if res.StatusCode == 404 {
		fmt.Printf("There is no %s directory in %s repository\n", directoryPath, repo.Name)
		return nil
	}
	bodyBytes := readBody(res)
	jsonBody := unmarshalArray(bodyBytes)
	dirContent := convertToContent(jsonBody)

	fmt.Printf("Found %d files in %s in %s repository\n", len(dirContent), directoryPath, repo.Name)
	return dirContent
}

func retrieveFileContent(fileContent *GHContent, headers map[string]string) string {
	res := callAPI(fileContent.Download_url, headers)
	if res.StatusCode == 404 {
		fmt.Printf("Error during retrieval of %s", fileContent.Path)
		fmt.Println(res)
		return ""
	}

	return string(readBody(res))
}

func convertToContent(jsonBody []map[string]any) []*GHContent {
	workflowsCount := len(jsonBody)
	output := make([]*GHContent, 0, workflowsCount)
	for _, jsonWorkflow := range jsonBody {
		workflow := newContent(jsonWorkflow)
		if workflow.Type == "file" {
			output = append(output, workflow)
		}
	}

	return output
}

func convertToRepos(jsonBody []map[string]any) []*Repository {
	reposCount := len(jsonBody)
	output := make([]*Repository, 0, reposCount)
	for _, jsonRepo := range jsonBody {
		// with type assertion
		output = append(output, newRepository(jsonRepo))
	}

	return output
}