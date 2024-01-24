package cntechkitgorequest

import (
	"bytes"
	"fmt"
	"net/http"
)

type collection struct {
	name     string
	baseUrl  string
	requests []*request
}

type CollectionOptions struct {
	BaseUrl string
}

func NewCollection(collectionName string, options CollectionOptions) *collection {
	if options.BaseUrl != "" {
		return &collection{
			name:    collectionName,
			baseUrl: options.BaseUrl,
		}
	}
	return &collection{
		name: collectionName,
	}
}

func (c *collection) AddRequest(r *request) *collection {
	c.requests = append(c.requests, r)
	return c
}

func (c *collection) CreateClients() *client {
	var _client client
	for _, r := range c.requests {
		var url string = r.url
		if c.baseUrl != "" {
			url = fmt.Sprintf("%v%v", c.baseUrl, r.url)
		}
		if r.path != "" {
			url = fmt.Sprintf("%v%v", r.url, r.path)
		}

		var body *bytes.Buffer
		if r.body != nil {
			body = bytes.NewBuffer(r.body)
		} else {
			body = bytes.NewBuffer([]byte{})
		}

		req, err := http.NewRequest(r.method, url, body)
		if err != nil {
			fmt.Println("error creating request for", r.key, "error: ", err)
		}
		query := req.URL.Query()
		for key, value := range r.queries {
			query.Add(key, value)
		}
		req.URL.RawQuery = query.Encode()

		for key, value := range r.headers {
			req.Header.Add(key, value)
		}

		if r.username != "" && r.password != "" {
			req.SetBasicAuth(r.username, r.password)
		}

		if _client == nil {
			_client = map[string]*http.Request{
				r.key: req,
			}
		} else {
			_client[r.key] = req
		}

	}
	return &_client

}

func (c *collection) Log() *collection {
	fmt.Println("Collection Name:\t", c.name)
	fmt.Println("Base Url:\t\t", c.baseUrl)
	fmt.Println("Request Count:\t\t", len(c.requests))
	for index, request := range c.requests {
		fmt.Println("\t*", index+1, "\t", "KEY:", request.key)
		fmt.Println("\t ", request.method, "\t", request.url)
		fmt.Println("\t ", "path:", request.path)
		fmt.Println("\t ", "headers:")
		for key, value := range request.headers {
			fmt.Println("\t \t", key, ":", value)
		}
		fmt.Println("\t ", "queries:")
		for key, value := range request.queries {
			fmt.Println("\t \t", key, ":", value)
		}
		fmt.Println("\t ", "body:", string(request.body))
		fmt.Println("\t ", "basic auth info:", request.username, request.password)
	}

	return c
}
