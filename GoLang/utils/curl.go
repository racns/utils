package utils

import (
	"encoding/json"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"strings"
)

type CurlResult struct {
	Data       any
	Body       string
	StatusCode int
	Header     http.Header
	Err        error
}

func CurlFunc(url string, method string, params map[string]any, headers map[string]any) (result CurlResult) {

	client := &http.Client{}
	body, _ := json.Marshal(params)
	request, _ := http.NewRequest(method, url, strings.NewReader(string(body)))

	if headers != nil {
		for key, val := range headers {
			request.Header.Set(key, cast.ToString(val))
		}
	}

	response, err := client.Do(request)
	if err != nil {
		result.Err = err
		return
	}
	result.StatusCode = response.StatusCode
	result.Header = response.Header

	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			result.Err = err
			return
		}
		result.Body = string(body)
		result.Data = Json.Decode(string(body))
	}

	return
}

func CurlGet(url string, params map[string]any, headers map[string]any) (result CurlResult) {

	request, _ := http.NewRequest("GET", url, nil)

	if params != nil {
		query := request.URL.Query()
		for key, val := range params {
			query.Add(key, cast.ToString(val))
		}
		request.URL.RawQuery = query.Encode()
	}

	if headers != nil {
		for key, val := range headers {
			request.Header.Set(key, cast.ToString(val))
		}
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		result.Err = err
		return
	}

	result.StatusCode = response.StatusCode
	result.Header = response.Header

	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			result.Err = err
			return
		}
		result.Body = string(body)
		result.Data = Json.Decode(string(body))
	}

	return
}

func CurlPost(url string, params map[string]any, headers map[string]any) (result CurlResult) {
	return CurlFunc(url, "POST", params, headers)
}

func CurlPut(url string, params map[string]any, headers map[string]any) (result CurlResult) {
	return CurlFunc(url, "PUT", params, headers)
}

func CurlDel(url string, params map[string]any, headers map[string]any) (result CurlResult) {
	return CurlFunc(url, "DELETE", params, headers)
}

func CurlPatch(url string, params map[string]any, headers map[string]any) (result CurlResult) {
	return CurlFunc(url, "PATCH", params, headers)
}