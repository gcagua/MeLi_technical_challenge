// implements a CLI tool to summarize text files using the huggingface API
package main

import (
	"fmt"
	"os"
)

// cli entry point
func main() {

	// validates the environment variables and arguments
	summaryType, text, token, err := ValidateEnvAndArgs()
	if err != nil {
		fmt.Println("An error occured while validating env variables and arguments", err)
		os.Exit(1)
	}

	// reads the file checking for forbidden routes
	content, err := ReadFileSecurely(text)
	if err != nil {
		fmt.Println("An error occured while reading the file", err)
		os.Exit(1)
	}

	// sanitizes the input file before passing it to the model
	sanitizedContent := SanitizeFile(content)

	// builds the prompt based on the summary type and the sanitized content
	prompt := BuildPrompt(summaryType, sanitizedContent)

	// makes the api request
	summary, err := MakeAPIRequest(prompt, token)
	if err != nil {
		fmt.Printf("Error making the api request: %v", err)
		os.Exit(1)
	}

	// prints the response
	fmt.Printf("%s\n", summary)
}
