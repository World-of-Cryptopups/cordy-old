package utils

import (
	"io/ioutil"
	"net/http"
)

// Fetcher is a custom fetcher for getting the json response.
func Fetcher(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return make([]byte, 0), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	return body, nil
}
