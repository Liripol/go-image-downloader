package downloader

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func DownloadImage(url, folder string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	if !strings.Contains(fileName, ".") {
		fileName += ".jpg"
	}

	os.MkdirAll(folder, os.ModePerm)
	filePath := path.Join(folder, fileName)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
