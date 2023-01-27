#ChatGPT

A simple command-line tool for interacting with the OpenAI GPT-3 API
What you'll need

    A valid OpenAI API key
    Go 1.16 or above

#How to Setup

    Clone this repository
    In the root directory of the repository, run go build to build the executable
    Run the executable with the following flags:
        -p for the prompt you want GPT-3 to respond to
        -k for your OpenAI API key

#Features

    Dynamic prompt handling:

The app accepts a prompt input from the command line using the -p flag.

    Automatic API key handling:

The app accepts an OpenAI API key from the command line using the -k flag.
#TODO

    Add support for multiple prompts in a single execution
    Implement error handling for invalid API keys
    Optimize the code for better performance

#How to use

    Run the executable with the -p flag followed by your prompt and the -k flag followed by your OpenAI API key.
    Example: ./chatgpt -p "What is the capital of France?" -k "your_api_key"
    The app will return the response from GPT-3.
