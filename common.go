package vlivego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiID = "8c6cc7b45d2568fb668be6e05b6e5a3b"

func sync(data interface{}, url string) interface{} {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
		decode(data, response.Body)
	}
	return data
}

func decode(data interface{}, body io.ReadCloser) interface{} {
	json.NewDecoder(body).Decode(data)
	return data
}
