package thingiverse

import (
	"fmt"
)

type DownloadLink struct {
	Filename    string `json:"name"`
	DownloadUrl string `json:"download_url"`
}

type Model struct {
	Name          string
	Token         string
	DownloadLinks []DownloadLink
}

func GetModel(modelId string) (Model, error) {
	var model Model
	html, err := getHtml(modelId)
	if err != nil {
		fmt.Println("could not fetch data from thingiverse homepage")
		return Model{}, err
	}
	javascriptBundleVersion, err := getJavascriptBundleVersion(html)
	if err != nil {
		fmt.Println("could not fetch bundle version")
		return Model{}, err
	}

	model.Name, err = getModelName(html)
	if err != nil {
		fmt.Println("could not get model name")
		return Model{}, err
	}
	javascript, err := getJavascriptBundle(javascriptBundleVersion)
	if err != nil {
		fmt.Println("could not fetch javascript")
		return Model{}, err
	}

	model.Token, err = getBearerToken(javascript)
	if err != nil {
		fmt.Println("could not get bearer token")
		return Model{}, err
	}

	model.DownloadLinks, err = getFiles(model.Token, modelId)
	if err != nil {
		fmt.Println("could not get download links")
		return Model{}, err
	}

	return model, nil
}
