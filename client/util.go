package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	DaemonEndp = "http://localhost:5555"
)

func CreateRequest(method, url string, body io.Reader) (*Response, error) {
	var response *Response

	req, err := http.NewRequest(method, DaemonEndp, body)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)

	return response, err

}
