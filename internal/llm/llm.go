package llm

import (
	"RepoLens/internal/config"
	"RepoLens/internal/prompts"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type OpenAIRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type OpenAIResponse struct {
	Output []struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	} `json:"output"`
}

func FixScript(cfg *config.Config, prevCode string, execError string) (string, error) {
	prompt := prompts.FixScriptPrompt(prevCode, execError)
	return callOpenAI(cfg, prompt)
}

func callOpenAI(cfg *config.Config, prompt string) (string, error) {
	reqBody := OpenAIRequest{
		Model: cfg.Model,
		Input: prompt,
	}
	jsonData, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/responses", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+cfg.OpenAIAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result OpenAIResponse
	json.NewDecoder(resp.Body).Decode(&result)

	if len(result.Output) == 0 {
		return "", fmt.Errorf("no response")
	}

	text := result.Output[0].Content[0].Text
	text = strings.ReplaceAll(text, "```go", "")
	text = strings.ReplaceAll(text, "```", "")
	text = strings.TrimSpace(text)

	return text, nil
}

func GenerateScript(cfg *config.Config, prompt string) (string, error) {
	return callOpenAI(cfg, prompt)
}

func AnalyzeRepo(cfg *config.Config, data string) (string, error) {
	prompt := prompts.ANALIZE_REPO + data
	return callOpenAI(cfg, prompt)
}
