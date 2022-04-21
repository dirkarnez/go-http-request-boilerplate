package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	values := map[string]interface{}{ /* */ }
	json_data, _ := json.Marshal(values)

	req, _ := http.NewRequest("POST", "https://reqbin.com/sample/post/json", bytes.NewBuffer(json_data))
	// req.Header.Set("mode", "cors")
	// req.Header.Set("credentials", "include")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp.StatusCode, string(body))
}
