package test

import (
	"encoding/json"
	"net/http"

	fastshot "github.com/opus-domini/fast-shot"
)

func NewClient() fastshot.ClientHttpMethods {
	builder := fastshot.NewClient("http://localhost:8001/api/v1")

	return builder.Build()
}

func GetJson(response *http.Response, target interface{}) error {
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}

// func Request(method string, path *string, data map[string]any) (*map[string]any, *error) {
// 	URL := "http://localhost:8001/api"

// 	var passData *bytes.Reader = nil

// 	if data != nil {
// 		buf, err := json.Marshal(data)

// 		if err != nil {
// 			return nil, &err
// 		}

// 		passData := bytes.NewReader(buf)

// 		// passData = &jsonData

// 		// reader := io.Read

// 		// passData = io.Reader.Read(&jsonData)
// 	}

// 	r, err := http.NewRequest(method, URL+*path, *passData)
// 	r.Header.Add("Content-Type", "application/json")

// }
