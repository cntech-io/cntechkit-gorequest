package cntechkitgorequest

import (
	"encoding/json"
	"fmt"
	"strings"
)

type request struct {
	method   string
	url      string
	key      string
	headers  map[string]string
	queries  map[string]string
	username string
	password string
	path     string
	body     []byte
}

func NewGetRequest(key string, url string) *request {
	return &request{
		method: "GET",
		url:    url,
		key:    key,
	}
}

func NewPostRequest(key string, url string) *request {
	return &request{
		method: "POST",
		url:    url,
		key:    key,
	}
}

func NewPutRequest(key string, url string) *request {
	return &request{
		method: "PUT",
		url:    url,
		key:    key,
	}
}

func NewPatchRequest(key string, url string) *request {
	return &request{
		method: "PATCH",
		url:    url,
		key:    key,
	}
}

func NewDeleteRequest(key string, url string) *request {
	return &request{
		method: "DELETE",
		url:    url,
		key:    key,
	}
}

func (r *request) AddHeader(key string, value string) *request {

	if r.headers == nil {
		r.headers = map[string]string{
			key: value,
		}
	} else {
		r.headers[key] = value
	}
	return r
}

func (r *request) AddBearerHeader(token string) *request {
	if r.headers == nil {
		r.headers = map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", token),
		}
	} else {
		r.headers["Authorization"] = fmt.Sprintf("Bearer %v", token)
	}
	return r
}

func (r *request) AddContentTypeJsonHeader() *request {
	if r.headers == nil {
		r.headers = map[string]string{
			"content-type": "application/json",
		}
	} else {
		r.headers["content-type"] = "application/json"
	}
	return r
}

func (r *request) AddQuery(key string, value string) *request {

	if r.queries == nil {
		r.queries = map[string]string{
			key: value,
		}
	} else {
		r.queries[key] = value
	}
	return r
}

func (r *request) AddPath(path string) *request {
	if strings.HasPrefix(path, "/") {
		r.path = path
	} else {
		r.path = fmt.Sprintf("/%v", path)
	}
	return r
}

func (r *request) AddBody(body interface{}) *request {
	bodyBtye, err := json.Marshal(body)
	if err != nil {
		fmt.Println("error marshalling body for", r.key, "error: ", err)
	}
	r.body = bodyBtye
	return r
}

func (r *request) AddBasicAuth(username string, password string) *request {
	r.username = username
	r.password = password
	return r
}
