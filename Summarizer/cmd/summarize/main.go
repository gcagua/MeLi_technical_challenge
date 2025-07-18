// implements a CLI tool to summarize text files using the huggingface API
package main

import (
	"fmt"
	"os"
	"time"
)

const (
	MaxFileSize    = 1024 * 1024 * 10 // 10MB limit,
	RequestTimeout = 30 * time.Second
)

// cli entry point
func main() {

	// validates the environment variables and arguments
	summaryType, text, token, err := validateEnvAndArgs()
	if err != nil {
		fmt.Println("An error occured while validating env variables and arguments", err)
		os.Exit(1)
	}

	// reads the file checking for forbidden routes
	content, err := readFileSecurely(text)
	if err != nil {
		fmt.Println("An error occured while reading the file", err)
		os.Exit(1)
	}

	// sanitizes the input file before passing it to the model
	sanitizedContent := sanitizeFile(content)

	// builds the prompt based on the summary type and the sanitized content
	prompt := buildPrompt(summaryType, sanitizedContent)

	// makes the api request
	summary, err := makeAPIRequest(prompt, token)
	if err != nil {
		fmt.Printf("Error making the api request: %v", err)
		os.Exit(1)
	}

	// prints the response
	fmt.Printf("%s\n", summary)
}

func buildPrompt(summaryType SummaryType, content string) string {
	return fmt.Sprintf("%s: \n\n%s", summaryType, content)
}

func stringTypeToEnum(summaryType string) (SummaryType, error) {
	switch summaryType {
	case "short":
		return Short, nil
	case "medium":
		return Medium, nil
	case "bullet":
		return Bullet, nil
	default:
		return Bullet, fmt.Errorf("the provided summary type is not valid")
	}
}
