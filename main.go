package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/atotto/clipboard"
)

var requestStruct = Format{GroupGUID: "[guid]", Domain: "bit.ly"}

func main() {

	text, _ := clipboard.ReadAll()
	fmt.Println(text)
	requestStruct.LongURL = text
	client := &http.Client{}
	jsonText, err := json.Marshal(requestStruct)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", "https://api-ssl.bitly.com/v4/shorten", bytes.NewBuffer(jsonText))
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", "Bearer "+"[token]")
	res, err := client.Do(request)
	response, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(response, &IdLinkStruct)
	if err != nil {
		panic(err)
	}
	clipboard.WriteAll(IdLinkStruct.ID)
}
