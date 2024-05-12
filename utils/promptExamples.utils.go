package utils

import (
	"encoding/json"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"os"
)

type PromptExample struct {
	Input1 string `json:"input1"`
	Input2 string `json:"input2"`
	Output string `json:"output"`
}

const HeaderInput1 = "Here is the company about us and careers page. Analyse it and generate a company profile."
const HeaderInput2 = "Here is a mail template for our institute, provide a mail template aligned to the companys profile, highlighting the benefits that they will have if they hire from our institue"
const HeaderOutput = "Provide a professional email, not more than 500 words, inviting the company to our recruit from our campus"

func GeneratePrompt(companyProfile string) []genai.Part {
	jsonFile, jsonOpenError := os.ReadFile("prompts.example.json")

	if jsonOpenError != nil {
		fmt.Println("Prompt Examples could not be loaded")
		print(jsonOpenError)
		return nil
	}

	var allExamples []PromptExample
	_ = json.Unmarshal(jsonFile, &allExamples)

	var examples []genai.Part

	for i := 0; i < len(allExamples); i++ {
		examples = append(examples, genai.Text(HeaderInput1+"\n"+allExamples[i].Input1))
		examples = append(examples, genai.Text(HeaderInput2+"\n"+allExamples[i].Input2))
		examples = append(examples, genai.Text(HeaderOutput+"\n"+allExamples[i].Output))
	}

	examples = append(examples, genai.Text(HeaderInput1+"\n"+companyProfile))
	examples = append(examples, genai.Text(HeaderInput2+"\n"+allExamples[0].Input2))
	examples = append(examples, genai.Text(HeaderOutput))

	return examples
}
