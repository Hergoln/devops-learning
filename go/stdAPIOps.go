package main

import (
	"os"
	"fmt"
	"time"
	"net/http"
)

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