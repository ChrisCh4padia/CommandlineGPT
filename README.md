# ChatGPT
A simple command-line tool for interacting with the OpenAI GPT-3 API

# What you'll need
- A valid OpenAI API key
- Go 1.16 or above

# How to Setup
1. Clone this repository
2. In the root directory of the repository, run `go build` to build the executable
3. Run the executable with the following flags: 
   - `-p` for the prompt you want GPT-3 to respond to
   - `-k` for your OpenAI API key

# Features
1. Dynamic prompt handling:

The app accepts a prompt input from the command line using the `-p` flag.

2. API key handling:

The app accepts an OpenAI API key from the command line using the `-k` flag.

# TODO
- Add support for multiple prompts in a single execution
- Implement error handling for invalid API keys

# How to use

1. Run the executable with the `-p` flag followed by your prompt and the `-k` flag followed by your OpenAI API key.
2. Example: `./chatgpt -p "What is the capital of France?" -k "this-is-your-api-key"`


