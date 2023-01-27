package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/solywsh/chatgpt"
)

// prompt is the user's input text
var prompt string

// response to request
var answer string

// API key for the GPTservice
var apikey string

var err error

func main() {
	// Define command-line flags for setting the prompt and API key
	flag.StringVar(&prompt, "p", "", "Your chatgpt prompt")
	flag.StringVar(&apikey, "k", "", "Your chatgpt API key")
	flag.Parse()

	// Check that the prompt is not empty
	if prompt != "" {
		generateGPTmessage()
		fmt.Println(answer)
	} else {
		fmt.Println("Prompt can not be empty!")
	}
}

func generateGPTmessage() {
	// The timeout is used to control the situation that the session is in a long and multi session situation.
	// If it is set to 0, there will be no timeout. Note that a single request still has a timeout setting of 30s.
	chat := chatgpt.New(apikey, "user_id(not required)", 0*time.Second)
	defer chat.Close()

	// Get the response
	answer, err = chat.Chat(prompt)
	if err != nil {
		if strings.Contains(err.Error(), "401") {
			fmt.Printf("Unauthorized. Please check your key")
		} else {
			fmt.Println(err)
		}

	}

}
