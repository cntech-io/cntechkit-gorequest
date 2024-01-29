package request

import (
	"bytes"
	"fmt"
	"net/http"
)

type Collection struct {
	name     string
	baseUrl  string
	requests []*request
}

type CollectionOptions struct {
	BaseUrl string
}

func NewCollection(collectionName string, options CollectionOptions) *Collection {
	if options.BaseUrl != "" {
		return &Collection{
			name:    collectionName,
			baseUrl: options.BaseUrl,
		}
	}
	return &Collection{
		name: collectionName,
	}
}

func (c *Collection) AddRequest(r *request) *Collection {
	c.requests = append(c.requests, r)
	return c
}

func (c *Collection) GetRequest(key string) *request {
	if len(c.requests) == 0 {
		return nil
	}
	for _, request := range c.requests {
		if request.key == key {
			return request
		}
	}
	return nil
}

func (c *Collection) UpdateRequest(key string, r *request) *Collection {
	if len(c.requests) == 0 {
		return c
	}

	for index, request := range c.requests {
		if request.key == key {
			c.requests[index] = r
		}
	}
	return c
}

func (c *Collection) CreateClient(key string) *Client {
	var _client Client
	var r *request = findRequest(c, key)

	var url string = makeUrl(c, r)

	var body *bytes.Buffer = makeBody(r)

	req, err := http.NewRequest(r.method, url, body)
	if err != nil {
		fmt.Println("error creating request for", r.key, "error: ", err)
	}

	req = AddQuery(req, r.queries)

	for key, value := range r.headers {
		req.Header.Add(key, value)
	}

	if r.username != "" && r.password != "" {
		req.SetBasicAuth(r.username, r.password)
	}

	_client = map[string]*http.Request{
		key: req,
	}

	return &_client
}

func (c *Collection) CreateClients() *Client {
	var _client Client
	for _, r := range c.requests {

		var url string = makeUrl(c, r)

		var body *bytes.Buffer = makeBody(r)

		req, err := http.NewRequest(r.method, url, body)
		if err != nil {
			fmt.Println("error creating request for", r.key, "error: ", err)
		}

		req = AddQuery(req, r.queries)

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

func (c *Collection) Log() *Collection {
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
