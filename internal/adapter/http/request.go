package infra

import (
	"io/ioutil"
	"net/http"
)

// MakeAPIRequest makes a request to a third-party API
func MakeAPIRequest(url string) ([]byte, error) {
	// Make the request
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
