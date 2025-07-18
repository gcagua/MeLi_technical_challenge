package sanitize

import (
	"fmt"
	"strings"
)

// sanitizes the files checking for the most common general and bart-specific prompt injections
func sanitizeFile(input string) string {
	promptInjectionPatterns := []string{
		"ignore previous instructions",
		"ignore all previous",
		"ignore everything",
		"do not summarize",
		"just write",
		"repeat this",
		"new instructions", // general prompt injections
		"TL;DR:",
		"IMPORTANT:",
		"### HEADLINE:",
		"Summary:",
		"## Summary:",
		"output",
		"Q:", // bart-specific prompt injections
	}

	// checks if the file contains any of the patterns
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
