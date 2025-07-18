// implements a CLI tool to summarize text files using the huggingface API
package main

import (
	"fmt"
	"os"

	"github.com/gcagua/MeLi_technical_challenge/Summarizer/internal/api"
	"github.com/gcagua/MeLi_technical_challenge/Summarizer/internal/file"
	"github.com/gcagua/MeLi_technical_challenge/Summarizer/internal/sanitize"
	"github.com/gcagua/MeLi_technical_challenge/Summarizer/internal/validation"
)

// cli entry point
func main() {

	// validates the environment variables and arguments
	summaryType, text, token, err := validation.ValidateEnvAndArgs()
	if err != nil {
		fmt.Println("An error occured while validating env variables and arguments", err)
		os.Exit(1)
	}

	// reads the file checking for forbidden routes
	content, err := file.ReadFileSecurely(text)
	if err != nil {
		fmt.Println("An error occured while reading the file", err)
		os.Exit(1)
	}

	// sanitizes the input file before passing it to the model
	sanitizedContent := sanitize.SanitizeFile(content)

	// builds the prompt based on the summary type and the sanitized content
	prompt := api.BuildPrompt(summaryType, sanitizedContent)

	// makes the api request
	summary, err := api.MakeAPIRequest(prompt, token)
	if err != nil {
		fmt.Printf("Error making the api request: %v", err)
		os.Exit(1)
	}

	// prints the response
	fmt.Printf("%s\n", summary)
}
