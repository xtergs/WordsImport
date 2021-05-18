package updates

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type gitHubResponse struct {
	TagName string   `json:"tag_name"`
	Assets  []assets `json:"assets"`
}

type assets struct {
	DownloadUrl string `json:"browser_download_url"`
	Name        string `json:"name"`
}

func CheckNewVersion(link string, currentVersion string) error {
	response, err := http.Get(link)
	if err != nil {
		return err
	}

	model := gitHubResponse{}
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil {
		return err
	}

	fmt.Printf("Latest version: %s, current: %s\n", model.TagName, currentVersion)

	if model.TagName > currentVersion {
		for _, asset := range model.Assets {
			platform := ""
			normalized := strings.ToLower(asset.Name)
			if strings.Contains(normalized, "linux") {
				platform = "Linux"
			}
			if strings.Contains(normalized, "windows") {
				platform = "Windows"
			}
			if strings.Contains(normalized, "darwin") {
				platform = "Mac"
			}
			fmt.Printf("%s - %s\n", platform, asset.DownloadUrl)
		}

	}

	return nil
}
