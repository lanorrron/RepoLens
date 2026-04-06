package runner

import (
	"RepoLens/internal/config"
	"RepoLens/internal/executor"
	"RepoLens/internal/llm"
	"RepoLens/internal/prompts"
	"fmt"
	"os"
)

func Log(msg string, info string) {
	fmt.Printf("%s: %s\n", msg, info)
}

func Exit(msg string, err error) {
	fmt.Printf("%s: %v\n", msg, err)
	os.Exit(1)
}

func GenerateScript(cfg *config.Config, lang string, repoPath string) string {
	fmt.Println("Generating analysis script...")

	prompt := fmt.Sprintf(prompts.ANALYSIS_SCRIPT, repoPath, lang)

	code, err := llm.GenerateScript(cfg, prompt)
	if err != nil {
		Exit("generate script error", err)
	}
	return code
}

func RunWithRetries(cfg *config.Config, code string, lang string, attempts int) string {
	var output string
	var err error

	for i := 1; i <= attempts; i++ {
		Log(fmt.Sprintf("Running script (attempt %d)", i), "")
		output, err = executor.Run(code)
		if err == nil {
			return output
		}

		fmt.Println("Script execution error:", err)
		fmt.Println("Output:", output)

		fmt.Println("Trying to fix script...")
		code, err = llm.FixScript(cfg, code, err.Error())
		if err != nil {
			Exit("fix error", err)
		}
	}

	Exit("failed after retries", err)
	return ""
}

func AnalyzeRepo(cfg *config.Config, data string) string {
	fmt.Println(" Analyzing repo...")
	report, err := llm.AnalyzeRepo(cfg, data)
	if err != nil {
		Exit("analysis error", err)
	}
	return report
}

func SaveReport(path string, report string) {
	if err := os.WriteFile(path, []byte(report), 0644); err != nil {
		Exit("write file error", err)
	}
}
