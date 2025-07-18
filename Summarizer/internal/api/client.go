package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// creates and executes the request to an api of the huggingface bart endpoint given a prompt and a token
func makeAPIRequest(prompt, token string) (string, error) {
	body := Request{Inputs: prompt}
	jsonData, err := json.Marshal(body)
	if err != nil { // check if there's an error marshalling the body
		return "", fmt.Errorf("Error marshaling the request: %v", err)
	}

	// creates request and adds headers
	req, err := http.NewRequest("POST", os.Getenv("HUGGINGFACE_ENDPOINT"), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("Error creating the request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: RequestTimeout,
	}

	// executes the request
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making the request: %v", err)
	}
	defer response.Body.Close()

	// reads the response from the output
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading the response: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Api error %d: %s", response.StatusCode, string(responseBody))
	}

	// parses the response into a Response object
	var parsedBody []Response
	if err := json.Unmarshal(responseBody, &parsedBody); err != nil {
		return "", fmt.Errorf("Error parsing the response: %v", err)
	}

	// checks if the response length is 0
	if len(parsedBody) == 0 {
		return "", fmt.Errorf("Api returned an empty response")
	}

	// checks if the response is an empty string
	if parsedBody[0].SummarizedText == "" {
		return "", fmt.Errorf("The response summary is empty")
	}

	// sanitizes the output file to remove any undesired patterns
	revisedText := sanitizeFile(parsedBody[0].SummarizedText)
	return revisedText, nil
}

// builds the prompt based on the summary type and the file content
func buildPrompt(summaryType SummaryType, content string) string {
	return fmt.Sprintf("%s: \n\n%s", summaryType, content)
}
