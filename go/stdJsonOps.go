package main

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