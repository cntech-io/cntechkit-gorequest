package main

import (
    "fmt"
    "strings"
    "net/http"
    "io"
    "encoding/json"
)

type client map[string]*http.Request

type collection struct {
    name string
    baseUrl string
    requests []*request
}

type CollectionOptions struct {
    baseUrl string
}

type request struct {
    method string
    url string
    key string
    headers map[string]string
    queries map[string]string
    username string
    password string
    path string
}


func (cl client) Result(key string) (interface{},error) {
    httpClient := &http.Client{}
    req := cl[key]
    resp ,err := httpClient.Do(req)
    if err!= nil {
        fmt.Println("error making request for",key,"error: ",err)
        return nil,err
    }
    defer resp.Body.Close()

    response, err :=io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("error reading response for",key,"error: ",err)
        return nil,err
    }
    var data interface{}
    err = json.Unmarshal(response, &data)
    if err != nil {
        fmt.Println("error unmaeshalling response json for",key,"error: ",err)
        return nil,err
    }
    return data,nil
}


func NewCollection(collectionName string,options CollectionOptions) *collection {
    if options.baseUrl!= "" {
        return &collection{
            name: collectionName,
            baseUrl: options.baseUrl,
        }
    }
    return &collection{
        name: collectionName,
    }
}

func (c *collection)AddRequest(r *request) *collection{
    c.requests = append(c.requests,r)
    return c
}

func (c *collection) CreateClients() *client {
    var  _client client
    for _, r := range c.requests{
        var url string = r.url
        if c.baseUrl != "" {
            url = fmt.Sprintf("%v%v",c.baseUrl,r.url)
        }
        if r.path != "" {
            url = fmt.Sprintf("%v%v",r.url,r.path)
        }


        req, err := http.NewRequest(r.method,url,nil)
        if err != nil {
            fmt.Println("error creating request for",r.key,"error: ",err)
        }
        query := req.URL.Query()
        for key, value := range r.queries {
            query.Add(key,value)
        }
        req.URL.RawQuery = query.Encode()

        for key, value := range r.headers {
            req.Header.Add(key, value)
        }

        if r.username != "" && r.password != "" {
            req.SetBasicAuth(r.username,r.password)
        }

        if _client == nil {
            _client = map[string]*http.Request{
                r.key : req,
            }
        }else{
            _client[r.key] = req
        }

    }
    return &_client

}

func  (c *collection) Log() *collection{
    fmt.Println("Collection Name:\t",c.name)
    fmt.Println("Base Url:\t\t", c.baseUrl)
    fmt.Println("Request Count:\t\t", len(c.requests))
    for index, request := range c.requests {
        fmt.Println("\t*",index+1,"\t","KEY:",request.key)
        fmt.Println("\t ",request.method,"\t",request.url)
        fmt.Println("\t ","path:",request.path)
        fmt.Println("\t ","headers:")
        for key,value := range request.headers {
            fmt.Println("\t \t",key, ":",value)
        }
        fmt.Println("\t ","queries:")
        for key,value := range request.queries {
            fmt.Println("\t \t",key, ":",value)
        }
        fmt.Println("\t ","basic auth info:",request.username,request.password)
    }

    return c
}

func NewGetRequest(key string,url string) *request {
    return &request{
        method: "GET",
        url: url,
        key: key,
    }
}

func NewPostRequest(key string,url string) *request {
    return &request{
        method: "POST",
        url: url,
        key: key,
    }
}

func NewPutRequest(key string,url string) *request {
    return &request{
        method: "PUT",
        url: url,
        key: key,
    }
}

func NewPatchRequest(key string,url string) *request {
    return &request{
        method: "PATCH",
        url: url,
        key: key,
    }
}

func NewDeleteRequest(key string,url string) *request {
    return &request{
        method: "DELETE",
        url: url,
        key: key,
    }
}


func (r *request)AddHeader(key string,value string) *request {

    if r.headers == nil {
        r.headers = map[string]string{
            key: value,
        }
    }else{
        r.headers[key]=value
    }
    return r
}

// TODO: add application json

func (r *request)AddBearerHeader(token string) *request {

    if r.headers == nil {
        r.headers = map[string]string{
            "Authorization": fmt.Sprintf("Bearer %v",token),
        }
    }else{
        r.headers["Authorization"]=fmt.Sprintf("Bearer %v",token)
    }
    return r
}

func (r *request)AddQuery(key string,value string) *request {

    if r.queries == nil {
        r.queries = map[string]string{
            key: value,
        }
    }else{
        r.queries[key]=value
    }
    return r
} 

func (r *request) AddPath(path string) *request {
    if strings.HasPrefix(path,"/") {
        r.path = path
    } else {
        r.path = fmt.Sprintf("/%v",path)
    }
    return r
}

func (r *request) AddBasicAuth(username string,password string)*request{
    r.username=username
    r.password=password
    return r
}

func main() {
    c:= NewCollection("trendyol",CollectionOptions{baseUrl:"https://api.chucknorris.io"}).
    AddRequest(
        NewGetRequest("first","/jokes/random"),
    ).
    Log().
    CreateClients()
    r,_ := c.Result("first")
    fmt.Println(r)
}
