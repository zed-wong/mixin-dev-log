package main

import(
	"log"
	"bytes"
        "io/ioutil"
	"encoding/json"

	"github.com/hashicorp/go-retryablehttp"
)

func HttpGet(url string)string{
	client := retryablehttp.NewClient()
	client.Logger = nil
        resp, err := client.Get(url)
        if err != nil{
		log.Println("HttpGet:",err)
        }
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil{
		log.Println("ioutil.ReadAll:",err)
        }
	resp.Body.Close()
	return string(body)
}

func HttpPost(url string, values map[string]string)string{
        jsonValue, _ := json.Marshal(values)
        resp, err := retryablehttp.Post(url, "application/json", bytes.NewBuffer(jsonValue))
        if err != nil{
                log.Println(err)
        }
        body, _ := ioutil.ReadAll(resp.Body)
        result := string(body)
	resp.Body.Close()
        return result
}

