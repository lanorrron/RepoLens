package llm

import (
	"RepoLens/internal/config"
	"RepoLens/internal/prompts"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
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

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/responses", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+cfg.OpenAIAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("openai bad status %d: %s", resp.StatusCode, strings.TrimSpace(string(bodyBytes)))
	}

	var result OpenAIResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return "", fmt.Errorf("decode response json: %w", err)
	}

	if len(result.Output) == 0 {
		return "", fmt.Errorf("no output in response")
	}

	var sb strings.Builder

	for _, out := range result.Output {
		for _, c := range out.Content {
			if strings.TrimSpace(c.Text) != "" {
				if sb.Len() > 0 {
					sb.WriteString("\n")
				}
				sb.WriteString(c.Text)
			}
		}
	}

	text := sb.String()
	text = strings.ReplaceAll(text, "```go", "")
	text = strings.ReplaceAll(text, "```", "")
	text = strings.TrimSpace(text)

	if text == "" {
		return "", fmt.Errorf("empty text content in response")
	}

	return text, nil
}

func GenerateScript(cfg *config.Config, prompt string) (string, error) {
	return callOpenAI(cfg, prompt)
}

func AnalyzeRepo(cfg *config.Config, data string) (string, error) {
	prompt := prompts.ANALIZE_REPO + data
	return callOpenAI(cfg, prompt)
}
