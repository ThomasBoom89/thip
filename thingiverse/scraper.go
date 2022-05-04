package thingiverse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func getHtml(modelId string) ([]byte, error) {
	url := fmt.Sprintf("https://www.thingiverse.com/thing:%s/files", modelId)
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return html, nil
}

func getJavascriptBundleVersion(html []byte) (string, error) {
	re, err := regexp.Compile(`js/app\.bundle\.js\?(\d+)`)
	if err != nil {
		return "", err
	}
	matches := re.FindSubmatch(html)

	return string(matches[1]), nil
}

func getModelName(html []byte) (string, error) {
	re, err := regexp.Compile(`>(.*)by`)
	if err != nil {
		return "", err
	}
	matches := re.FindSubmatch(html)
	re, err = regexp.Compile(`[^\.A-Za-z0-9]+`)
	match := strings.TrimSpace(string(matches[1]))
	name := re.ReplaceAllString(match, "_")

	return name, nil
}

func getJavascriptBundle(version string) ([]byte, error) {
	url := fmt.Sprintf("https://cdn.thingiverse.com/site/js/app.bundle.js?%s", version)
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	javascript, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return javascript, nil
}

func getBearerToken(javascript []byte) (string, error) {
	re, err := regexp.Compile(`u="([a-z0-9]*)"`)
	if err != nil {
		return "", err
	}
	matches := re.FindSubmatch(javascript)

	return string(matches[1]), nil
}

func getFiles(token, modelId string) ([]DownloadLink, error) {
	var downloadLinks []DownloadLink
	url := fmt.Sprintf("https://api.thingiverse.com/things/%s/files", modelId)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&downloadLinks)
	if err != nil {
		return nil, err
	}

	return downloadLinks, nil
}
