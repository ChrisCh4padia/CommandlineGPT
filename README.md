# ChatGPT
A simple command-line tool for interacting with the OpenAI GPT-3 API

## What you need
- A valid OpenAI API key
- Go 1.16 or above

## How to Setup
1. Clone repo
2. Run `go build` to build the executable
3. Run the executable with the following flags: 
   - `-p` for the prompt you want GPT to respond to
   - `-k` for your OpenAI API key
   - `-h` Get some help!
   
   
## Using Environment variables
   You can also set the API key through the `gptkey` environment variable.
   
   Something like: 
   
   `export gptkey=YOURAPIKEY`

   This allows you to send requests without the `-k` parameter.
   
## Features
1. Dynamic prompt handling:

The app accepts a prompt input from the command line using the `-p` flag.

2. API key handling:

The app accepts an OpenAI API key from the command line using the `-k` flag or through the `gptkey` environment variable.

## TODO
- Add support for multiple prompts in a single execution
- Implement error handling for invalid API keys

## How to use

With implemented environment variable:

![image](https://user-images.githubusercontent.com/87771733/218195459-1571877e-be99-42fc-9be9-83642c88a4de.png)



With `-k` parameter:

![image](https://user-images.githubusercontent.com/87771733/218196298-29b94f23-e8ca-4bdb-a667-98b06404f624.png)

