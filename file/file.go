package file

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(filePath, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return nil
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
