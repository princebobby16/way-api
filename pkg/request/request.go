package request

import (
	"io"
	"net/http"
)

func Request(url string, data io.Reader, method string, header, query map[string]string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	if len(header) >= 1 {
		for key, value := range header {
			request.Header.Set(key, value)
		}
	}
	if len(query) >= 1 {
		requestQuery := request.URL.Query()
		for key, value := range query {
			requestQuery.Add(key, value)
		}
		request.URL.RawQuery = requestQuery.Encode()
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
