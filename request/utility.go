package request

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func makeUrl(c *Collection, r *request) string {
	var url string = r.url
	if c.baseUrl != "" {
		if strings.HasPrefix(r.url, "/") {
			url = fmt.Sprintf("%v%v", c.baseUrl, r.url)
		} else {
			url = fmt.Sprintf("%v/%v", c.baseUrl, r.url)
		}
	}
	if r.path != "" {
		if strings.HasPrefix(r.path, "/") {
			url = fmt.Sprintf("%v%v", r.url, r.path)
		} else {
			url = fmt.Sprintf("%v/%v", r.url, r.path)
		}
	}
	return url
}

func makeBody(r *request) *bytes.Buffer {
	var body *bytes.Buffer
	if r.body != nil {
		body = bytes.NewBuffer(r.body)
	} else {
		body = bytes.NewBuffer([]byte{})
	}
	return body
}

func AddQuery(req *http.Request, queries map[string]string) *http.Request {
	q := req.URL.Query()
	for key, value := range queries {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	return req
}

func findRequest(c *Collection, key string) *request {
	for _, request := range c.requests {
		if request.key == key {
			return request
		}
	}
	return nil
}
