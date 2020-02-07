package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type asd struct {
	Current string `json:"current"`
}

func main() {
	url := "http://35.193.162.111/run/frequent"
	fmt.Println("URL:>", url)
	current_time := time.Now()
	t := current_time.Format("15:04:05")
	time := &asd{}
	current := fmt.Sprintf("%s", t)
	time.Current = current
	fmt.Println(time.Current)
	jsonStr, _ := json.Marshal(time)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("PostReq: ", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
