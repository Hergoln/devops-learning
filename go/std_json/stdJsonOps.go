package std_json

import (
	"encoding/json"
	"fmt"
	"io"
	"my_module/std_errors"
	"net/http"
	"os"
)

func UnmarshalArray(body []byte) []map[string]any {
	var jsonBody []map[string]any

	err := json.Unmarshal(body, &jsonBody)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(std_errors.UNMARSHAL_ERROR)
	}

	return jsonBody
}

func ReadBody(response *http.Response) []byte {
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)

	if response.StatusCode == http.StatusOK {
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(std_errors.BYTE_READ_ERROR)
		}
		return bytes
	}

	fmt.Println("No body found")
	return nil
}
