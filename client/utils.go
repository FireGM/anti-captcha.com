package client

import (
	"io"
	"net/http"
	"os"
)

// saveFromUrl taked url to image and return path of img on
func saveFromUrl(url string) (path string, err error) {
	path = "captcha.jpg"
	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(path)
	if err != nil {
		return
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return
	}
	err = file.Close()
	if err != nil {
		return
	}
	return
}
