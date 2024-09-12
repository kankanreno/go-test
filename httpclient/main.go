package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	curDateStr := time.Now().Format("2006-01-02")
	url := fmt.Sprintf("https://api.sjtu.edu.cn/v1/enterprise/calendar?from=%s&to=%s", curDateStr, curDateStr)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Json Decoding
	// {"errno":0,"error":"Succeed.","total":0,"entities":[{"day":"2024-09-12","type":5}]}
	type Entitie struct {
		Day  string `json:"day"`
		Type int8   `json:"type"`
	}
	var result struct {
		Errno    int8      `json:"errno"`
		Error    string    `json:"error"`
		Total    int8      `json:"total"`
		Entities []Entitie `json:"entities"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
	if result.Errno != 0 {
		fmt.Println(result.Error)
	}

	var ret = struct {
		ServerTime string `json:"serverTime"`
		Type       int8   `json:"type"`
	}{
		time.Now().Format("2006-01-02 15:04:05"),
		result.Entities[0].Type,
	}
	fmt.Println(ret)
}
