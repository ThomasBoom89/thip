package thingiverse

import (
	"archive/zip"
	"io/ioutil"
	"net/http"
	"os"
)

func ToZipArchive(model Model) (string, error) {
	client := http.Client{}
	filename := model.Name + ".zip"
	file, err := os.Create("./" + filename)
	z := zip.NewWriter(file)
	for _, downloadLink := range model.DownloadLinks {
		req, err := http.NewRequest("GET", downloadLink.DownloadUrl, nil)
		if err != nil {
			return "", err
		}
		req.Header.Add("Authorization", "Bearer "+model.Token)
		response, err := client.Do(req)
		if err != nil {
			return "", err
		}
		file, err := z.Create(downloadLink.Filename)
		if err != nil {
			return "", err
		}
		body, err := ioutil.ReadAll(response.Body)
		_, err = file.Write(body)
		if err != nil {
			return "", err
		}
	}
	err = z.Close()
	if err != nil {
		return "", err
	}

	return filename, nil
}
