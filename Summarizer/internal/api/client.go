package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gcagua/MeLi_technical_challenge/Summarizer/internal/sanitize"
	"github.com/gcagua/MeLi_technical_challenge/Summarizer/types"
)

const (
	RequestTimeout = 30 * time.Second
)

// creates and executes the request to an api of the huggingface bart endpoint given a prompt and a token
func MakeAPIRequest(prompt, token string) (string, error) {
	body := types.Request{Inputs: prompt}
	jsonData, err := json.Marshal(body)
	if err != nil { // check if there's an error marshalling the body
		return "", fmt.Errorf("error marshaling the request: %v", err)
	}

	// creates request and adds headers
	// this is the documentation for the huggingface api: https://huggingface.co/facebook/bart-large-cnn
	req, err := http.NewRequest("POST", os.Getenv("HUGGINGFACE_ENDPOINT"), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating the request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: RequestTimeout,
	}

	// executes the request
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making the request: %v", err)
	}
	defer response.Body.Close()

	// reads the response from the output
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api error %d: %s", response.StatusCode, string(responseBody))
	}

	// parses the response into a Response object
	var parsedBody []types.Response
	if err := json.Unmarshal(responseBody, &parsedBody); err != nil {
		return "", fmt.Errorf("error parsing the response: %v", err)
	}

	// checks if the response length is 0
	if len(parsedBody) == 0 {
		return "", fmt.Errorf("api returned an empty response")
	}

	// checks if the response is an empty string
	if parsedBody[0].SummarizedText == "" {
		return "", fmt.Errorf("the response summary is empty")
	}

	// sanitizes the output file to remove any undesired patterns
	revisedText := sanitize.SanitizeFile(parsedBody[0].SummarizedText)
	return revisedText, nil
}

// builds the prompt based on the summary type and the file content
func BuildPrompt(summaryType types.SummaryType, content string) string {
	return fmt.Sprintf("%s: \n\n%s", summaryType, content)
}
