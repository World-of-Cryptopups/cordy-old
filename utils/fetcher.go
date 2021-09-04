package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

	defer resp.Body.Close()

	return body, nil
}

// PostFetcher is Fetcher but with POST data.
func PostFetcher(data interface{}, url string) ([]byte, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return make([]byte, 0), err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(d))
	if err != nil {
		return make([]byte, 0), err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		fmt.Println(string(body))
		return make([]byte, 0), errors.New("user not found")
	}

	return body, nil
}
