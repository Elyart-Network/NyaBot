package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetRequest(url string, params string) ([]byte, error) {
	Request, err := http.Get(url + params)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(Request.Body)
	Context, err := io.ReadAll(Request.Body)
	if err != nil {
		return nil, err
	}
	return Context, nil
}

func PostRequest(url string, params any) ([]byte, error) {
	ByteSlice, err := json.Marshal(params)
	Request, err := http.Post(url, "application/json", bytes.NewBuffer(ByteSlice))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(Request.Body)
	Context, err := io.ReadAll(Request.Body)
	if err != nil {
		return nil, err
	}
	return Context, nil
}
