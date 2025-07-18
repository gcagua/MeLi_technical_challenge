package validation

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gcagua/MeLi_technical_challenge/Summarizer/types"
)

// parses and validates CLI glags and environment variables
func ValidateEnvAndArgs() (types.SummaryType, string, string, error) {
	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		return types.Short, "", "", fmt.Errorf("huggingface token not found")
	}

	typeParam := flag.String("type", "", "Type of the summary")
	shortTypeParam := flag.String("t", "", "Short type of the summary")
	inputParam := flag.String("input", "", "The input article")
	flag.Parse()

	var typeValue string // parses the summary-type value
	if *typeParam != "" {
		typeValue = *typeParam
	} else if *shortTypeParam != "" {
		typeValue = *shortTypeParam
	}

	var inputValue string // parses the input-file name value
	if *inputParam != "" {
		inputValue = *inputParam
	} else {
		args := flag.Args()
		if len(args) > 0 {
			inputValue = args[0]
		}
	}

	if typeValue == "" { // if summary-type value is an empty string, returns the error
		return types.Short, "", "", fmt.Errorf("summary type is required (--type or -t)")
	}
	typeValue = strings.ToLower(typeValue)

	enumValue, err := types.StringTypeToEnum(typeValue) // checks if the summary type exists
	if err != nil {
		return types.Short, "", "", err
	}

	if inputValue == "" { // if the name of the file is an empty strng, returns the error
		return types.Short, "", "", fmt.Errorf("input type is required (--input or as an argument)")
	}

	return enumValue, inputValue, token, nil
}
