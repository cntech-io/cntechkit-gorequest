package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type client map[string]*http.Request

func (cl client) Result(key string) (interface{}, error) {
	httpClient := &http.Client{}
	req := cl[key]
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("error making request for", key, "error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response for", key, "error: ", err)
		return nil, err
	}
	var data interface{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		fmt.Println("error unmarshalling response json for", key, "error: ", err)
		return nil, err
	}
	return data, nil
}
