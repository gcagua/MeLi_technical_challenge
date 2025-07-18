package main

// ¿Hay algún requerimiento del idioma en el que se debe hacer el resumen? ✓ inglés
// ¿Se puede enviar más de un archivo por solucion? ✓ creo que no importa
// ¿El codigo debe ser sensible a inputs en mayusculas o minusculas? ✓ no
// ¿Que debería pasar si el input del tipo de resumen está mal ingresado? ¿debería haber un tipo de resumen por defecto en ese caso? ✓ no debería funcionar
// ___________________________________________________________________________
// Si alguno de los parámetros necesarios no está, simplemente saca el error
// ___________________________________________________________________________
// ¿El archivo tiene un tamaño máximo?

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type SummaryType int

const (
	Short SummaryType = iota
	Bullet
	Medium
)

var summaryTypePrompt = map[SummaryType]string{
	Short:  "a concise summary (min 1 and max 2 sentences)",
	Bullet: "a list of bullet points",
	Medium: "a paragraph summary",
}

func (summaryType SummaryType) String() string {
	return summaryTypePrompt[summaryType]
}

// ---------------------------------------------------------

type Request struct {
	Inputs string `json:"inputs"`
}

type Response struct {
	SummarizedText string `json:"summary_text"`
}

// --------------------------------------------------------------------------
func main() {

	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		fmt.Println("Huggingface token not found")
		os.Exit(1)
	}

	typeParam := flag.String("type", "", "Type of the summary")
	shortTypeParam := flag.String("t", "", "Short type of the summary")
	inputParam := flag.String("input", "", "The input article")
	flag.Parse()

	var typeValue string
	if *typeParam != "" {
		typeValue = *typeParam
	} else if *shortTypeParam != "" {
		typeValue = *shortTypeParam
	}

	var inputValue string
	if *inputParam != "" {
		inputValue = *inputParam
	} else {
		args := flag.Args()
		if len(args) > 0 {
			inputValue = args[0]
		}
	}

	if typeValue == "" {
		fmt.Println("Summary type is required (--type or -t)")
		os.Exit(1)
	}
	typeValue = strings.ToLower(typeValue)

	if _, err := stringTypeToEnum(typeValue); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	if inputValue == "" {
		fmt.Println("Input type is required (--input or as an argument)")
		os.Exit(1)
	}

	// siguiente: se debe validar que los inputs sean válidos. {
	// * si alguno de los inputs no existe -> error  ✓
	// * si ambos existen pero el tipo de summary no está dentro de los permitidos -> error ✓
	// * si el parámetro del summary se ingresa mal e.g. --t bullet -> error (revisar si queda tiempo) ✓
	// * si ambos existen pero el archivo no tiene una finalización en txt -> error ✓
	// * si ambos existen pero el archivo no existe en el directorio -> error ✓
	//}

	content, err := readFileSecurely(inputValue)
	if err != nil {
		fmt.Println("An error occured while reading the file", err)
		os.Exit(1)
	}

	summaryType, _ := stringTypeToEnum(typeValue)
	sanitizedContent := sanitizeInputFile(content)

	prompt := fmt.Sprintf("Summarize this in the style of %s: \n\n%s", summaryType, sanitizedContent)
	// body := Request{Inputs: prompt}
	// jsonData, err := json.Marshal(body)

	// req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/facebook/bart-large-cnn", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	fmt.Println("An error occured with the api request", err.Error())
	// 	return
	// }

	// req.Header.Add("Authorization", "Bearer "+token)
	// req.Header.Add("Content-Type", "application/json")

	// client := &http.Client{}
	// response, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("An error occured while making request", err.Error())
	// 	return
	// }
	// defer response.Body.Close()

	// responseBody, _ := io.ReadAll(response.Body)
	// var parsedBody []Response
	// if err := json.Unmarshal(responseBody, &parsedBody); err != nil {
	// 	fmt.Println("ehehehe new error", err.Error())
	// 	return
	// }

	summary, err := makeAPIRequest(prompt, token)
	if err != nil {
		fmt.Printf("Error making the api request: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", summary)
}

func sanitizeInputFile(input string) string {
	promptInjectionPatterns := []string{
		"ignore previous instructions",
		"ignore all previous",
		"ignore everything",
		"new instructions",
		"TL;DR:",
		"IMPORTANT:",
		"### HEADLINE:",
		"Summary:",
	}

	input = strings.ToLower(input)
	for _, pattern := range promptInjectionPatterns {
		pattern_lower := strings.ToLower(pattern)
		if strings.Contains(input, pattern_lower) {
			fmt.Println("The input file contains malicious content")
			input = strings.ReplaceAll(input, pattern, "")
		}
	}
	return input
}

func makeAPIRequest(prompt, token string) (string, error) {
	body := Request{Inputs: prompt}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("Error marshaling the request: %v", err)
	}

	req, err := http.NewRequest("POST", os.Getenv("HUGGINGFACE_ENDPOINT"), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("Error creating the request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making the request: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading the response: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Api error %d: %s", response.StatusCode, string(responseBody))
	}

	var parsedBody []Response
	if err := json.Unmarshal(responseBody, &parsedBody); err != nil {
		return "", fmt.Errorf("Error parsing the response: %v", err)
	}

	if len(parsedBody) == 0 {
		return "", fmt.Errorf("Api returned an empty response")
	}

	if parsedBody[0].SummarizedText == "" {
		return "", fmt.Errorf("The response summary is empty")
	}

	return parsedBody[0].SummarizedText, nil
}

func readFileSecurely(path string) (string, error) {
	cleanPath := filepath.Clean(path)
	if strings.Contains(cleanPath, "..") {
		return "", fmt.Errorf("Path traversal not allowed")
	}

	if !strings.HasSuffix(cleanPath, ".txt") {
		return "", fmt.Errorf("File is not of type txt")
	}

	fileInfo, err := os.Stat(cleanPath)
	if err != nil {
		return "", fmt.Errorf("File was not found: %v", err)
	}

	if !fileInfo.Mode().IsRegular() {
		return "", fmt.Errorf("Not a regular file")
	}

	content, err := os.ReadFile(cleanPath)
	if err != nil {
		return "", fmt.Errorf("File could not be read: %v", err)
	}

	return string(content), nil
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
