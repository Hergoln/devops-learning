// github structs
package main


type Repository struct {
	Name string
	Url string
	ContentsUrl string
}

func newRepository(jsonRepo map[string]any) *Repository {
	repo := Repository{
		Name: jsonRepo["full_name"].(string),
		Url: jsonRepo["url"].(string),
		ContentsUrl: jsonRepo["contents_url"].(string),
	}
	return &repo
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

func newContent(jsonContent map[string]any) *GHContent {
  workflow := GHContent{
		// with type assertion
		Name: 				jsonContent["name"].(string),
		Path: 				jsonContent["path"].(string),
		Url: 					jsonContent["url"].(string),
		Git_url: 			jsonContent["git_url"].(string),
		Download_url: jsonContent["download_url"].(string),
		Type: 				jsonContent["type"].(string),
	}
	if jsonContent["type"] == "file" {
		if jsonContent["content"] != nil {
			workflow.Content = jsonContent["content"].(string)
		}
	}
	return &workflow
}