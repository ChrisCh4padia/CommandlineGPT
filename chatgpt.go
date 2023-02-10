package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var prompt string
var answer string
var apikey string
var err error

func main() {

	flag.StringVar(&prompt, "p", "", "Your chatgpt prompt")
	flag.StringVar(&apikey, "k", "", "Your chatgpt API key")
	flag.Parse()

	if apikey == "" {
		envvar := os.Getenv("gptkey")
		if envvar == "" {
			fmt.Println("Error: No API key specified. Please set the key through the -k flag or the gptkey environment variable")
			return
		} else {
			apikey = envvar
		}
	}

	if prompt != "" {
		generateGPTmessage()
		fmt.Println(answer)
	} else {
		fmt.Println("Prompt can not be empty!")
	}
}

func generateGPTmessage() {

	body := map[string]interface{}{
		"prompt":     prompt,
		"model":      "text-davinci-003",
		"max_tokens": 100,
	}
	jsonBody, _ := json.Marshal(body)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(jsonBody))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apikey))

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	var response map[string]interface{}
	json.Unmarshal(respBody, &response)

	choices := response["choices"].([]interface{})
	choice := choices[0].(map[string]interface{})
	text := choice["text"].(string)

	text = strings.TrimSpace(text)

	fmt.Println(text)
}
