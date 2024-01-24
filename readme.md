## Methods

| Method                                                          | Description                                              |
| --------------------------------------------------------------- | -------------------------------------------------------- |
| NewCollection(collectionName string, options CollectionOptions) | Creates collection                                       |
| collection.AddRequest(r Request)                                | Adds request to collection                               |
| collection.CreateClients()                                      | Creates client                                           |
| client.Result(key string)                                       | Make request to added request with key                   |
| NewGetRequest(key string, url string)                           | Creates GET request                                      |
| NewPostRequest(key string, url string)                          | Creates POST request                                     |
| NewPatchRequest(key string, url string)                         | Creates PATCH request                                    |
| NewPutRequest(key string, url string)                           | Creates PUT request                                      |
| NewDeleteRequest(key string, url string)                        | Creates DELETE request                                   |
| request.AddHeader(key string, value string)                     | Adds header to request                                   |
| request.AddBearerHeader(token string)                           | Adds Bearer header to request                            |
| request.AddContentTypeJsonHeader()                              | Adds "content-type":"application/json" header to request |
| request.AddQuery(key string, value string)                      | Adds query to request                                    |
| request.AddPath(path string)                                    | Adds path to request                                     |
| request.AddBody(body interface{})                               | Adds body to request                                     |
| request.AddBasicAuth(username string, password string)          | Adds basic auth to request                               |
