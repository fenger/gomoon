package file

import (
	"fmt"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	filepath := "./index"
	url := "http://github.com"
	fmt.Println("downloading", url)
	DownloadFile(filepath, url)
	fmt.Printf("downloaded %s to %s\n", url, filepath)
}
