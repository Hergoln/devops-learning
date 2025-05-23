package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
)

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
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		body = string(bodyBytes)
	}

	fmt.Printf("Body: %s\n", body)
}