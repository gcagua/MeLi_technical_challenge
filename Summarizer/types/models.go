package types

import "fmt"

type SummaryType int

const (
	Short SummaryType = iota
	Bullet
	Medium
)

var summaryTypePrompt = map[SummaryType]string{
	Short:  "Make a summary (1-2 sentences) of: ",
	Bullet: "Make a list of bullet points of: ",
	Medium: "Write a paragraph summary of the following: ",
}

func (summaryType SummaryType) String() string {
	return summaryTypePrompt[summaryType]
}

type Request struct {
	Inputs string `json:"inputs"`
}

type Response struct {
	SummarizedText string `json:"summary_text"`
}

func StringTypeToEnum(summaryType string) (SummaryType, error) {
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
