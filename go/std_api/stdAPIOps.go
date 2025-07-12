package std_api

import (
	"fmt"
	"my_module/std_errors"
	"net/http"
	"os"
	"time"
)

func CallAPI(url string, headers map[string]string) *http.Response {
	fmt.Println("Calling url: ", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(std_errors.REQUEST_ERROR)
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
		os.Exit(std_errors.API_CALL_ERROR)
	}

	return res
}
